package exceptions

import (
	"fmt"
)

func LoginError(err error) {
	fmt.Printf("%s ### Error while try login ### \n==============================\n%s\n==============================\n%s",
		red, err, reset)
}

func RegistrationError(err error) {
	fmt.Printf("%s ### Error while try register ### \n==============================\n%s\n==============================\n%s",
		red, err, reset)
}

func UnauthorisedError() {
	fmt.Printf("%s ### Unauthorised error ###\n%s", red, reset)
}
