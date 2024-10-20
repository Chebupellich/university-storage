package main

import (
	"client/requests"

	"client/services"
)

func main() {
	requests.CheckAuth()

	for {
		if !requests.IsAuth {
			services.StartAuth()
			continue
		}

		services.RunUserService()
	}
}
