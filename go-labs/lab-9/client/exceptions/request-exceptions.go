package exceptions

import (
	"fmt"
)

const red = "\033[31m"
const reset = "\033[0m"

func CreateRequestError(err error) {
	fmt.Printf("%s ### Create request error ### \n==============================\n%s\n==============================\n%s",
		red, err, reset)
}

func BadRequestError(err error) {
	fmt.Printf("%s ### Bad request error ### \n==============================\n%s\n==============================\n%s",
		red, err, reset)
}

func InternalServerError(err error) {
	fmt.Printf("%s ### Internal server error ### \n==============================\n%s\n==============================\n%s",
		red, err, reset)
}

func GeneralError(err error) {
	fmt.Printf("%s ### Response error ### \n==============================\n%s\n==============================\n%s",
		red, err, reset)
}

func GetRequestError(err error) {
	fmt.Printf("%s ### Error while get response ### \n==============================\n%s\n==============================\n%s",
		red, err, reset)
}
