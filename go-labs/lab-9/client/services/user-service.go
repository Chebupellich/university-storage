package services

import (
	"client/requests"

	"client/utils"
	"fmt"
	"strconv"

	tm "github.com/buger/goterm"
)

func RunUserService() {
	for {
		fmt.Println("----- Bloba app -----")
		fmt.Println("# Choose option:")

		fmt.Println("1. Get user by name")
		fmt.Println("2. Get all users")
		fmt.Println("3. Create user")
		fmt.Println("4. Update user")
		fmt.Println("5. Delete user")

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
			default:
				fmt.Println(tm.Color("\tIncorrect input: type number of option", tm.RED))
				continue
			}
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
			fmt.Println(tm.Color("\tUser not found", tm.RED))
			continue
		}
		fmt.Printf("User: \n%v", user)

		if scanner.Scan() {
			utils.CallClear()
			break
		} else {
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
			fmt.Println(tm.Color("\tUsers not found", tm.RED))
			break
		}

		fmt.Println("Users:")
		for val := range resp {
			fmt.Println(val)
		}

		if scanner.Scan() {
			utils.CallClear()
			break
		} else {
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
			age, err := strconv.Atoi(scanner.Text())
			if err != nil || age < 1 || age > 200 {
				fmt.Println(tm.Color("\tIncorrect input: type correct integer age", tm.RED))
				continue
			}
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
			fmt.Println(tm.Color("\tUser not created", tm.RED))
			break
		}
		fmt.Println("User has been created")

		if scanner.Scan() {
			utils.CallClear()
			break
		} else {
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
			age, err := strconv.Atoi(scanner.Text())
			if err != nil || age < 1 || age > 200 {
				fmt.Println(tm.Color("\tIncorrect input: type correct integer age", tm.RED))
				continue
			}
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
			continue
		}
		fmt.Printf("User: \n%v", user)

		if scanner.Scan() {
			utils.CallClear()
			break
		} else {
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

		utils.CallClear()

		if scanner.Scan() {
			utils.CallClear()
			break
		} else {
			utils.CallClear()
			break
		}
	}
}
