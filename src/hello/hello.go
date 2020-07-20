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

const countMonitoring = 4
const delayMonitoring = 5
const fileLogger = "log.txt"
const fileSitesForMonitoring = "log.txt"

func main() {
	for {
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			monitoring()
		case 2:
			printLogs()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("command not found")
			os.Exit(-1)
		}
	}
}

func monitoring() {
	sites := readFiles()

	for i := 0; i < countMonitoring; i++ {
		for _, site := range sites {

			isOnlineHealthSite(site)
		}
		fmt.Println("await next monitoring .. ")
		time.Sleep(delayMonitoring * time.Second) //scheduler monitoring

	}
	fmt.Println(" ")
}
func isOnlineHealthSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode == 200 {
		registerLogger(site, true)
	} else {
		registerLogger(site, false)
	}
}

func readFiles() []string {
	var sites []string

	file, err := os.Open(fileSitesForMonitoring)

	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF { //End Of File
			break
		}
	}

	file.Close()
	return sites
}

func showMenu() {
	fmt.Println("input option list")
	fmt.Println("1 - Init Monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}

func readCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)

	return commandRead
}

func registerLogger(site string, status bool) {

	file, err := os.OpenFile(fileLogger, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " is online " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLogs() {

	file, err := ioutil.ReadFile(fileLogger) // mudar para constante

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
}
