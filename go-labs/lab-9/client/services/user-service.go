package services

import (
	"client/exceptions"
	"client/requests"

	"client/utils"
	"fmt"
	"strconv"

	tm "github.com/buger/goterm"
)

func RunUserService() {
	for {
		if !requests.IsAuth {
			exceptions.UnauthorisedError()

			return
		}

		fmt.Println("----- Bloba app -----")
		fmt.Println("# Choose option:")

		fmt.Println("1. Get user by name")
		fmt.Println("2. Get all users")
		fmt.Println("3. Create user")
		fmt.Println("4. Update user")
		fmt.Println("5. Delete user")
		fmt.Println("6. Get admin access")
		fmt.Println("7. Logout")

		fmt.Print("\n> ")
		if scanner.Scan() {
			val, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(tm.Color("\tIncorrect input: type number of option", tm.RED))
				continue
			}

			switch val {
			case 1:
				utils.CallClear()
				GetUserHandler()
			case 2:
				utils.CallClear()
				UsersHandler()
			case 3:
				utils.CallClear()
				CreateUserHandler()
			case 4:
				utils.CallClear()
				UpdateUserHandler()
			case 5:
				utils.CallClear()
				DeleteUserHandler()
			case 6:
				utils.CallClear()
				RootHandler()
			case 7:
				utils.CallClear()
				Logout()
			default:
				fmt.Println(tm.Color("\tIncorrect input: type number of option", tm.RED))
				continue
			}
		}
	}
}
func Logout() {
	var logout string

	fmt.Println("----- Bloba app -----")
	for {
		fmt.Println("# Type /exit to go back")
		fmt.Println("# Type Y to logout:")
		fmt.Print("> ")

		if scanner.Scan() {
			text := scanner.Text()
			if text == "/exit" {
				utils.CallClear()
				break
			}
			logout = text
		} else {
			fmt.Println(tm.Color("\tIncorrect input", tm.RED))
			continue
		}

		if len(logout) == 0 {
			fmt.Println(tm.Color("\tIncorrect input: not allowed empty strings", tm.RED))
			continue
		} else if logout == "Y" {
			continue
		}

		requests.Logout()
		fmt.Println("--- You have been logout ---")
		if scanner.Scan() {
			utils.CallClear()
			break
		}
	}
}

func GetAdminAccess() {
	var password string

	fmt.Println("----- Bloba app -----")
	for {
		fmt.Println("# Type /exit to go back")
		fmt.Println("# Type admin password:")
		fmt.Print("> ")

		if scanner.Scan() {
			text := scanner.Text()
			if text == "/exit" {
				utils.CallClear()
				break
			}
			password = text
		} else {
			fmt.Println(tm.Color("\tIncorrect input", tm.RED))
			continue
		}

		if len(password) == 0 {
			fmt.Println(tm.Color("\tIncorrect input: not allowed empty strings", tm.RED))
			continue
		}

		requests.AdminAccess(password)

		if scanner.Scan() {
			utils.CallClear()
			break
		}
	}
}

func GetUserHandler() {
	var username string

	fmt.Println("----- Bloba app -----")
	for {
		fmt.Println("# Type /exit to go back")
		fmt.Println("# Type user name:")
		fmt.Print("> ")

		if scanner.Scan() {
			text := scanner.Text()
			if text == "/exit" {
				utils.CallClear()
				break
			}
			username = text
		} else {
			fmt.Println(tm.Color("\tIncorrect input", tm.RED))
			continue
		}

		if len(username) == 0 {
			fmt.Println(tm.Color("\tIncorrect input: not allowed empty strings", tm.RED))
			continue
		}

		user := requests.GetUser(username)
		if user == nil {
			break
		}
		fmt.Printf("User: \n%v", user)

		if scanner.Scan() {
			utils.CallClear()
			break
		}
	}
}

func UsersHandler() {
	utils.CallClear()

	fmt.Println("----- Bloba app -----")
	for {
		resp := requests.GetUsers()
		if resp == nil {
			break
		}

		fmt.Println("Users:")
		for _, val := range resp {
			fmt.Printf("Name: %s\tAge: %d\n", val.Name, val.Age)
		}

		if scanner.Scan() {
			utils.CallClear()
			break
		}
	}
}

func CreateUserHandler() {
	var username string
	var age int

	fmt.Println("----- Bloba app -----")

	for {
		fmt.Println("# Type /exit to go back")
		fmt.Println("# Type username:")
		fmt.Print("> ")

		if scanner.Scan() {
			text := scanner.Text()
			if text == "/exit" {
				utils.CallClear()
				break
			}
			username = text
		} else {
			fmt.Println(tm.Color("\tIncorrect input", tm.RED))
			continue
		}

		fmt.Println("# Type age:")
		fmt.Print("> ")

		if scanner.Scan() {
			iage, err := strconv.Atoi(scanner.Text())
			if err != nil || iage < 1 || iage > 200 {
				fmt.Println(tm.Color("\tIncorrect input: type correct integer age", tm.RED))
				continue
			}
			age = iage
		} else {
			fmt.Println(tm.Color("\tIncorrect input", tm.RED))
			continue
		}

		if len(username) == 0 {
			fmt.Println(tm.Color("\tIncorrect input: not allowed empty strings", tm.RED))
			continue
		}

		res := requests.CreateUser(username, age)
		if !res {
			exceptions.InternalServerError(nil)
			break
		}
		fmt.Println("User has been created")

		if scanner.Scan() {
			utils.CallClear()
			break
		}
	}
}

func UpdateUserHandler() {
	var username string
	var newUsername string
	var age int

	fmt.Println("----- Bloba app -----")

	for {
		fmt.Println("# Type /exit to go back")
		fmt.Println("# Type CURRENT username:")
		fmt.Print("> ")

		if scanner.Scan() {
			text := scanner.Text()
			if text == "/exit" {
				utils.CallClear()
				break
			}
			username = text
		} else {
			fmt.Println(tm.Color("\tIncorrect input", tm.RED))
			continue
		}

		fmt.Println("# Type NEW username:")
		fmt.Print("> ")

		if scanner.Scan() {
			newUsername = scanner.Text()
		} else {
			fmt.Println(tm.Color("\tIncorrect input", tm.RED))
			continue
		}

		fmt.Println("# Type NEW age:")
		fmt.Print("> ")

		if scanner.Scan() {
			iage, err := strconv.Atoi(scanner.Text())
			if err != nil || iage < 1 || iage > 200 {
				fmt.Println(tm.Color("\tIncorrect input: type correct integer age", tm.RED))
				continue
			}
			age = iage
		} else {
			fmt.Println(tm.Color("\tIncorrect input", tm.RED))
			continue
		}

		if len(username) == 0 {
			fmt.Println(tm.Color("\tIncorrect input: not allowed empty strings", tm.RED))
			continue
		}

		user := requests.UpdateUser(username, newUsername, age)
		if user == nil {
			fmt.Println(tm.Color("\tUser not found", tm.RED))
			break
		}
		fmt.Printf("User: \n%v", user)

		if scanner.Scan() {
			utils.CallClear()
			break
		}
	}
}

func DeleteUserHandler() {
	var username string

	fmt.Println("----- Bloba app -----")
	for {
		fmt.Println("# Type /exit to go back")
		fmt.Println("# Type username:")
		fmt.Print("> ")

		if scanner.Scan() {
			text := scanner.Text()
			if text == "/exit" {
				utils.CallClear()
				break
			}
			username = text
		} else {
			fmt.Println(tm.Color("\tIncorrect input", tm.RED))
			continue
		}

		if len(username) == 0 {
			fmt.Println(tm.Color("\tIncorrect input: not allowed empty strings", tm.RED))
			continue
		}

		res := requests.DeleteUser(username)
		if !res {
			fmt.Println(tm.Color("\tUser not deleted", tm.RED))
			break
		}
		fmt.Println(tm.Color("\tUser deleted", tm.RED))

		if scanner.Scan() {
			utils.CallClear()
			break
		}
	}
}

func RootHandler() {
	var passwd string

	fmt.Println("----- Bloba app -----")
	for {
		fmt.Println("# Type /exit to go back")
		fmt.Println("# Type root password:")
		fmt.Print("> ")

		if scanner.Scan() {
			text := scanner.Text()
			if text == "/exit" {
				utils.CallClear()
				break
			}
			passwd = text
		} else {
			fmt.Println(tm.Color("\tIncorrect input", tm.RED))
			continue
		}

		if len(passwd) == 0 {
			fmt.Println(tm.Color("\tIncorrect input: not allowed empty strings", tm.RED))
			continue
		}

		res := requests.GetAdminAccess(passwd)
		if !res {
			fmt.Println(tm.Color("\tRoot not access", tm.RED))
			break
		}

		if scanner.Scan() {
			utils.CallClear()
			break
		}
	}
}
