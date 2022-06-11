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
		// Also you can do this instead (like most go-lib pkgs):
		// return e.Msg + ": " + e.Err.Error()
	}
	return e.Msg
}

func (e *CustomError) Unwrap() error {
	return e.Err
}

func main() {
	level1Err := errors.New("L1-Boom")
	level2Err := fmt.Errorf("L2-Ouch: %w", level1Err) // <-- Remember: fmt.Errorf() NOT adds a ':' by it´s own!
	level3Err := &CustomError{"L3-Wank", level2Err}
	level4Err := fmt.Errorf("L4-Toot: %w", level3Err) // <-- Remember: fmt.Errorf() NOT adds a ':' by it´s own!
	printAllWrappedErrors(level4Err)
	printCustomErrorOnly(level4Err)
	printCustomErrorIncludingAllWrappedErrors(level4Err)
	// Console output:
	// Print1 --> L4-Toot: L3-Wank: L2-Ouch: L1-Boom
	// Print2 --> L3-Wank
	// Print3 --> L3-Wank: L2-Ouch: L1-Boom
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
