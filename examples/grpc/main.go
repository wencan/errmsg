package main

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
		logger.Error("failed to servc", zap.Error(err))
	}
}
