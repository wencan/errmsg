package main

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
