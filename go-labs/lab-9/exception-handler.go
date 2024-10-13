package main

import "fmt"

func IncorrectValueError(err error) {
	fmt.Errorf("\033[31m* Incorrect value error: %v\033[31m", err)
}

func BadRequestError(err error) {
	fmt.Errorf("\033[31m* Bad request error: %v\033[31m", err)
}

func SendRequestError(err error) {
	fmt.Errorf("\033[31m* Can't sand request: %v\033[31m", err)
}
