package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const countMonitoring = 4
const schedullerTimeLoopMonitoring = 5

func main() {

	showIntroduction()

	for {
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			monitoring()
		case 2:
			fmt.Println("Logs ..")
		case 0:
			fmt.Println("Exit .. ")
		default:
			fmt.Println("command not found")
			os.Exit(-1)
		}
	}

}

func showIntroduction() {
	name := "Douglas"
	version := 1.1

	fmt.Println("hello, sr ", name)
	fmt.Println("program is version ", version)
	fmt.Println(" ")
	fmt.Println("input option list")
}

func showMenu() {
	fmt.Println("1 - Init Monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}

func readCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)

	return commandRead
}

func monitoring() {
	sites := []string{"http://www.alura.com.br", "http://www.google.com.br",
		"https://random-status-code.herokuapp.com/"}

	for i := 0; i < countMonitoring; i++ {
		for _, site := range sites {
			http.Get(site)

			isOlineHealthSite(site)
		}
		fmt.Println("await next monitoring .. ")
		time.Sleep(schedullerTimeLoopMonitoring * time.Second) //scheduler monitoring

	}

	fmt.Println(" ")
}

func isOlineHealthSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("website", site, "is online")
	} else {
		fmt.Println("website", site, "is offline")
	}
}

//TODO: colocar os sites em um arquivo txt e utilizar leitura de arquivo
