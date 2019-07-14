package errmsg_test

/*
 * Testable example
 *
 * wencan
 * 2019-07-13
 */

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/wencan/errmsg"
)

func ExampleWrapError() {
	errmsg.SetFlags(errmsg.FstdFlag | errmsg.Fshortfile)
	defer errmsg.SetFlags(errmsg.FstdFlag)

	getError := func() error {
		return errors.New("this is a test")
	}

	doSomeThing := func() error {
		err := getError()
		if err != nil {
			// Wrap error
			return errmsg.WrapError(errmsg.ErrUnavailable, err)
		}
		return nil
	}

	err := doSomeThing()
	fmt.Println("Error:", err)

	data, e := json.Marshal(err)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("Json:", string(data))

	errMsg := err.(*errmsg.ErrMsg)
	fmt.Println("File:", errMsg.File)
	fmt.Println("Line:", errMsg.Line)

	// Output:
	// Error: this is a test
	// Json: {"message":"this is a test","status":"Unavailable"}
	// File: example_test.go
	// Line: 30
}
