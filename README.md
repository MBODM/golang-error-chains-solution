# golang-error-chains-solution
Simple sample code, for a StackOverflow question about error-chains in Go.

#### StackOverflow question:
https://stackoverflow.com/questions/72505935/fmt-println-stops-printing-chain-at-wrapped-custom-error-golang

#### Solution:
The solution, how to do some flexible handling of error-chains containg a custom-error, is:
```golang
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
	}
	return e.Msg
}

func (e *CustomError) Unwrap() error {
	return e.Err
}
```

This makes it possible, when using wrapped errors in an error-chain, to filter out a single custom-error at top (in `main()` function) and show this custom-error's msg, while at the same time not loosing all wrapped error messages, that all shows up when doing a `fmt.Println(topLevelErr)` at the top:
```golang
func printAllWrappedErrors(topLevelError error) {
	fmt.Println(topLevelError)
}

func printCustomErrorOnly(topLevelError error) {
	var e *CustomError
	if errors.As(topLevelError, &e) {
		fmt.Println(e.Msg) // <-- This is the difference
	}
}

func printCustomErrorIncludingAllWrappedErrors(topLevelError error) {
	var e *CustomError
	if errors.As(topLevelError, &e) {
		fmt.Println(e) // <-- This is the difference
	}
}
```

#### Requirements:
You need at least Go 1.13 (released September 2019), when the error-chains/error-wrapping feature was added.

##### Have fun.
