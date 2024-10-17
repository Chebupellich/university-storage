package exceptions

import (
	"fmt"
)

func ParseError(err error) {
	fmt.Printf("%s ### Error while try parse ### \n==============================\n%s\n==============================\n%s",
		red, err, reset)
}

func CannotReadFileError(err error) {
	fmt.Printf("%s ### Error while read file ### \n==============================\n%s\n==============================\n%s",
		red, err, reset)
}
