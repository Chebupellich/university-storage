package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Print("Текущее время: ", currentTime.Format("15:04:05 2006-01-02"))
}
