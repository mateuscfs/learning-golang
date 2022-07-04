package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	viewIntroduction()

	for {
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
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
}

func viewIntroduction() {
	name, age := getNameAndAge()
	version := 1.1
	fmt.Println("Hello Mr.", name, ", your age is", age)
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

func getNameAndAge() (string, int) {
	name := "Mateus"
	age := 28

	return name, age
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	sites := getSites()

	for i := 0; i < len(sites); i++ {
		res, _ := http.Get(sites[i])

		if res.StatusCode == 200 {
			fmt.Println("Site: ", sites[i], "is online!")
		} else {
			fmt.Println("Site: ", sites[i], "is offline!")
		}
	}
}

func getSites() []string {
	sites := []string{"https://www.google.com", "https://www.alura.com.br"}
	sites = append(sites, "https://www.caelum.com.br")

	return sites
}
