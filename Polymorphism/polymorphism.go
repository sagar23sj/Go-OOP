// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {

// 	err := errors.New("it's an error")

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }

package main

import (
	"fmt"
)

// The error built-in interface type is the conventional interface
// for representing an error condition, with the nil value
// representing no error.

// the following interface type is already present in builtin package
// of golang source code, so even if you remove following interface
// this file will run without error
type error interface {
	Error() string
}

// errorString is a trivial implementation of error.
// implementation copied from errors package
type errorString struct {
	s string
}

//Error method from errors package
func (e *errorString) Error() string {
	return e.s
}

// MyError type for customised error message
//Final error message will contain severity as well as message
type MyError struct {
	ErrorSeverity string
	ErrorMessage  string
}

//Error method for MyError type
func (err *MyError) Error() string {
	return fmt.Sprintf("Error_Severity: %v    Error_Message: %s", err.ErrorSeverity, err.ErrorMessage)
}

func main() {

	//error object of erroString type
	err1 := &errorString{s: "Default Error function called from errors package"}

	//error object of MyError type
	err2 := &MyError{
		ErrorSeverity: "FATAL",
		ErrorMessage:  "Customised Error Function called from current package with error severity",
	}

	//General error -
	// can be used in normal error cases
	printError(err1)

	//Error with severity level -
	//can be used when there is need to log errors
	printError(err2)
}

//printError function with error interface parameter
//following function will only accpet those parameters which satisfy error interface
func printError(errMessage error) {
	//following line performs runtime polymorphism
	//i.e based on underlying type of object
	//compiler will decide which  Error method to execute
	fmt.Printf("\n %s", errMessage.Error())
}
