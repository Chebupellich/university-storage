package main

import (
	"bufio"
	"client/requests"
	"client/utils"
	"fmt"
	"os"
	"strconv"

	tm "github.com/buger/goterm"
)

var scanner = bufio.NewScanner(os.Stdin)

func StartAuth() {
	fmt.Println("----- Authentication -----")
	fmt.Println("# Choose option:")

	fmt.Println("1. Log in")
	fmt.Println("2. Sign in")

	for {
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
				Login()
			case 2:
				utils.CallClear()
				Register()
			default:
				fmt.Println(tm.Color("\tIncorrect input: type number of option", tm.RED))
				continue
			}
			break
		}
		break
	}
}

func Login() {
	var nick string
	var passwd string

	fmt.Println("----- Login -----")

	for {
		fmt.Println("# Type your nickname:")
		fmt.Print("> ")

		if scanner.Scan() {
			nick = scanner.Text()
		} else {
			fmt.Println(tm.Color("\tIncorrect input", tm.RED))
			continue
		}

		fmt.Println("# Type password:")
		fmt.Print("> ")

		if scanner.Scan() {
			passwd = scanner.Text()
		} else {
			fmt.Println(tm.Color("\tIncorrect input", tm.RED))
			continue
		}

		if len(nick) == 0 || len(passwd) == 0 {
			fmt.Println(tm.Color("\tIncorrect input: not allowed empty strings", tm.RED))
			continue
		}

		resp, err := requests.Login(nick, passwd)
		if err != nil {
			break
		}
		fmt.Printf("[RESPONSE]	%s", resp.Body)
		utils.CallClear()

		break
	}
}

func Register() {
	var nick string
	var passwd string

	fmt.Println("----- Register -----")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("# Type your nickname:")
		fmt.Print("> ")

		if scanner.Scan() {
			nick = scanner.Text()
		} else {
			fmt.Println(tm.Color("\tIncorrect input", tm.RED))
			continue
		}

		fmt.Println("# Type password:")
		fmt.Print("> ")

		if scanner.Scan() {
			passwd = scanner.Text()
		} else {
			fmt.Println(tm.Color("\tIncorrect input", tm.RED))
			continue
		}

		if len(nick) == 0 || len(passwd) == 0 {
			fmt.Println(tm.Color("\tIncorrect input: not allowed empty strings", tm.RED))
			continue
		}

		resp, err := requests.Register(nick, passwd)
		if err != nil {
			break
		}
		fmt.Printf("[RESPONSE]	%s", resp.Body)
		utils.CallClear()

		break
	}
}
