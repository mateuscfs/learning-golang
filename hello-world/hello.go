package main

import (
	"fmt"
)

func main() {
	name := "Mateus"
	version := 1.1
	fmt.Println("Hello Mr.", name)
	fmt.Println("This program is on version", version)

	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit Program")

	var command int
	fmt.Scanf("%d", &command)

	fmt.Println("Command choose was", command, "and it's address is", &command)
}
