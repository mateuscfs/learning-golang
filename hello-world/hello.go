package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitorSize = 3
const monitorDelay = 5
const dateFormat = "02/01/2006 15:04:05 -07:00"

func main() {
	viewIntroduction()

	for {
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			readLogs()
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
	fmt.Println((""))
	return readCommand
}

func getNameAndAge() (string, int) {
	name := "Mateus"
	age := 28

	return name, age
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	sites := readSites()

	for i := 0; i < monitorSize; i++ {
		for _, site := range sites {
			res, err := http.Get(site)

			printErr(err)

			if res.StatusCode == 200 {
				fmt.Println("Site: ", site, "is online!")
				log(site, true)
			} else {
				fmt.Println("Site: ", site, "is offline!")
				log(site, false)
			}
		}
		time.Sleep(monitorDelay * time.Second)
		fmt.Println("")
	}
}

func readSites() []string {
	file, err := os.Open("sites.txt")

	printErr(err)

	var sites []string

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func printErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func log(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	printErr(err)

	file.WriteString(site + " - online: " + strconv.FormatBool(status) + " | " + time.Now().Format(dateFormat) + "\n")

	file.Close()
}

func readLogs() {
	file, err := ioutil.ReadFile("log.txt")

	printErr(err)

	fmt.Println(string(file))
}
