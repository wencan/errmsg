package errmsg_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

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
		// output: ERROR   http/main.go:36 doSomeThing fail        {"status": "Unavailable", "message": "this is a test", "file": "main.go", "line": 21}
		handler.Logger.Error("doSomeThing fail", errmsg_zap.Fields(err)...)

		w.WriteHeader(errmsg.HTTPStatus(err))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		// write response
		// body: {"status":"Unavailable","message":"this is a test"}
		data, _ := json.Marshal(err)
		w.Write(data)
	} else {
		w.Write([]byte("{\"message\": \"hello, world\"}"))
	}
}

func ExampleHTTP_zap() {
	errmsg.SetFlags(errmsg.FstdFlag | errmsg.Fshortfile)

	logger := zap.NewExample() // It's so thoughtful!
	logger = logger.WithOptions(zap.AddStacktrace(zap.DPanicLevel))

	handler := &Handler{
		Logger: logger,
	}

	server := httptest.NewServer(handler)
	defer server.Close()

	res, err := http.Get(server.URL)
	if err != nil {
		logger.Error("GET fail", zap.Error(err))
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error("GET fail", zap.Error(err))
		return
	}
	logger.Info(res.Status, zap.Int("statusCode", res.StatusCode), zap.String("body", string(body)))

	// Output:
	// {"level":"error","msg":"doSomeThing fail","status":"Unavailable","message":"this is a test","file":"example_http_test.go","line":23}
	// {"level":"info","msg":"503 Service Unavailable","statusCode":503,"body":"{\"message\":\"this is a test\",\"status\":\"Unavailable\"}"}
}
