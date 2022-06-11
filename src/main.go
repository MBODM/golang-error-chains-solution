package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Msg string
	Err error
}

func (e *CustomError) Error() string {
	if e.Err != nil {
                // Added a ':' char here, because fmt.Errorf() does NOT
                // automatically add this (as i accidentally thought) !
		wrappedError := fmt.Errorf("%s: %w", e.Msg, e.Err)
		wrappedErrorMsg := wrappedError.Error()
		return wrappedErrorMsg
		// Also you can do this instead (as most golib stuff do):
		// return e.Msg + ": " + e.Err.Error()
	}
	return e.Msg
}

func (e *CustomError) Unwrap() error {
	return e.Err
}

func main() {
	level1Err := errors.New("[Error in L1]: Boom")
	level2Err := fmt.Errorf("[Error in L2]: Wrap L1Err %w", level1Err)
	level3Err := &CustomError{"[Error in L3]: Wrap L2Err", level2Err}
	level4Err := fmt.Errorf("[Error in L4]: Wrap L3Err %w", level3Err)
	printAllWrappedErrors(level4Err)
	printCustomErrorOnly(level4Err)
	printCustomErrorIncludingAllWrappedErrors(level4Err)
	// Console output:
	// Print1 --> [Error in L4]: Wrap L3Err [Error in L3]: Wrap L2Err [Error in L2]: Wrap L1Err [Error in L1]: Boom
	// Print2 --> [Error in L3]: Wrap L2Err
	// Print3 --> [Error in L3]: Wrap L2Err [Error in L2]: Wrap L1Err [Error in L1]: Boom
}

func printAllWrappedErrors(topLevelError error) {
	fmt.Println("Print1 -->", topLevelError)
}

func printCustomErrorOnly(topLevelError error) {
	var e *CustomError
	if errors.As(topLevelError, &e) {
		fmt.Println("Print2 -->", e.Msg) // <-- This is the difference
	}
}

func printCustomErrorIncludingAllWrappedErrors(topLevelError error) {
	var e *CustomError
	if errors.As(topLevelError, &e) {
		fmt.Println("Print3 -->", e) // <-- This is the difference
	}
}
