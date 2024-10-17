package main

import (
	"client/requests"

	"client/services"
)

func main() {
	requests.CheckAuth()
	services.StartAuth()

	services.RunUserService()
}
