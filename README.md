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

```go
	errmsg.SetFlags(errmsg.FstdFlag | errmsg.Fshortfile)

	err := errors.New("this is a test")
	err = errmsg.Wrap(errmsg.ErrUnavailable, err)

	// Output: this is a test
	fmt.Println(err.Error())

	// Output: status: Unavailable, message: this is a test, file:???.go, line: ???
	fmt.Println(err.String())

	// Output: {"status":"Unavailable","message":"this is a test"}
	s := json.Marshal(err)

	// Output: 503
	fmt.Println(errmsg.HTTPStatus(err))

	// Output: Unavailable
	fmt.Println(errmsg_grpc.Code(err))
	// Output: rpc error: code = Unavailable desc = this is a test
	fmt.Println(errmsg_grpc.Status(err).Err())

	zapLogger := zap.NewExample()
	// Output: ERROR   ???.go:?? error information        {"status": "Unavailable", "message": "this is a test", "file": "???.go", "line": ??}
	zapLogger.Error("error information", errmsg_zap.Fields(err)...)

	logrusLogger := logrus.New()
	// Output: time="... ..." level=error msg="error information" file=???.go line=?? message="this is a test" status=Unavailable
	logrusLogger.WithFields(errmsg_logrus.Fields(err)).Error("error information")
```