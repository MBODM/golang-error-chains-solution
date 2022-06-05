# golang-error-chains-solution
Simple sample code, for a StackOverflow question about error-chains in Go.

#### StackOverflow question:
https://stackoverflow.com/questions/72505935/fmt-println-stops-printing-chain-at-wrapped-custom-error-golang

#### Solution:
The solution, how to handle error-chains containg a custom error, is
```golang
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
```

This makes it possible, when using wrapped errors in an error-chain, to filter out a single custom error at top (in `main()` function) and show this custom error's msg, while at the same time not loosing all wrapped error messages, that shows up when doing a `fmt.Println(topLevelErr)` at the top.
