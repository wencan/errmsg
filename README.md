# errmsg
structured error representation

## Feature
* structured error
* classify error base on status
* Built-in 17 error status
* translate error status to HTTP status and grpc codes, etc
* translate error information to JSON body and grpc status, etc
* wrap location where the error occurred
* integrated zap and logrus

## Usage

### HTTP Server

```go
import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/wencan/errmsg"
	errmsg_zap "github.com/wencan/errmsg/logging/zap"
	"go.uber.org/zap"
)

func getError() error {
	return errors.New("this is a test")
}

func doSomeThing() error {
	err := getError()
	if err != nil {
		// Wrap error
		return errmsg.WrapError(errmsg.ErrUnavailable, err)
	}
	return nil
}

type Handler struct {
	Logger *zap.Logger
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := doSomeThing()
	if err != nil {
		// log
		// output: ERROR   http/main.go:36 doSomeThing fail        {"status": "Unavailable", "message": "this is a test", "file": "main.go", "line": 22}
		handler.Logger.Error("doSomeThing fail", errmsg_zap.Fields(err)...)

		// write response
		// body: {"status":"Unavailable","message":"this is a test"}
		data, _ := json.Marshal(err)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(data)
	} else {
		w.Write([]byte("{\"message\": \"hello, world\"}"))
	}
}

func main() {
	errmsg.SetFlags(errmsg.FstdFlag | errmsg.Fshortfile)

	logger, _ := zap.NewDevelopment()

	handler := &Handler{
		Logger: logger,
	}
	err := http.ListenAndServe("127.0.0.1:8080", handler)
	if err != nil && err != http.ErrServerClosed {
		logger.Error("failed to serve", zap.Error(err))
	}
}
```

### GRPC Server

```go
import (
	"context"
	"errors"
	"net"

	"github.com/wencan/errmsg"
	errmsg_grpc "github.com/wencan/errmsg/grpc"
	errmsg_zap "github.com/wencan/errmsg/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/reflection"
)

func getError() error {
	return errors.New("this is a test")
}

func doSomeThing() error {
	err := getError()
	if err != nil {
		// Wrap error
		return errmsg.WrapError(errmsg.ErrUnavailable, err)
	}
	return nil
}

type server struct {
	Logger *zap.Logger
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	err := doSomeThing()
	if err != nil {
		// log
		// output: ERROR   grpc/main.go:40 doSomeThing fail        {"status": "Unavailable", "message": "this is a test", "file": "main.go", "line": 25}
		s.Logger.Error("doSomeThing fail", errmsg_zap.Fields(err)...)

		// return grpc/status.Status
		// response:
		//	 ERROR:
		//	  Code: Unavailable
		//	  Message: this is a test
		return nil, errmsg_grpc.Status(err).Err()
	}

	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	errmsg.SetFlags(errmsg.FstdFlag | errmsg.Fshortfile)

	logger, _ := zap.NewDevelopment()

	ln, err := net.Listen("tcp", ":8087")
	if err != nil {
		logger.Error("listen fail", zap.Error(err))
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{Logger: logger})
	reflection.Register(s) // for grpcurl
	if err := s.Serve(ln); err != nil {
		logger.Error("failed to serve", zap.Error(err))
	}
}
```