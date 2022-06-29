package main

import (
	"fmt"
	"os"
)

func main() {
	viewIntroduction()

	showMenu()

	command := readCommand()

	switch command {
	case 1:
		fmt.Println("Monitor")
	case 2:
		fmt.Println("Logs")
	case 0:
		fmt.Println("Bye Bye")
		os.Exit(0)
	default:
		fmt.Println("Option inserted don't exist")
		os.Exit(-1)
	}
}

func viewIntroduction() {
	name := "Mateus"
	version := 1.1
	fmt.Println("Hello Mr.", name)
	fmt.Println("This program is on version", version)
}

func showMenu() {
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit Program")
}

func readCommand() int {
	var readCommand int
	fmt.Scan(&readCommand)

	fmt.Println("Command choose was", readCommand, "and it's address is", &readCommand)
	return readCommand
}
