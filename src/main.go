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
		wrappedError := fmt.Errorf("%s %w", e.Msg, e.Err)
		wrappedErrorMsg := wrappedError.Error()
		return wrappedErrorMsg
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
	fmt.Println("----------")
	printCustomError(level4Err)
	fmt.Println("----------")
}

func printAllWrappedErrors(topLevelError error) {
	fmt.Println(topLevelError)
}

func printCustomError(topLevelError error) {
	var e *CustomError
	if errors.As(topLevelError, &e) {
		fmt.Println(e.Msg) // <-- This is the difference
	}
}

func printCustomErrorIncludingAllBelow(topLevelError error) {
	var e *CustomError
	if errors.As(topLevelError, &e) {
		fmt.Println(e) // <-- This is the difference
	}
}
