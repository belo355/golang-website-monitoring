package main

import (
	"fmt"
	"net/http"
	"os"
)

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
			os.Exit(0)
		case 0:
			fmt.Println("Exit .. ")
			os.Exit(0)
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
	var sites [4]string
	sites[0] = "http://www.alura.com.br"
	sites[1] = "http://www.google.com.br"
	sites[2] = ""
	sites[3] = ""

	site := "http://www.alura.com.br"
	fmt.Println("watch status in ..", site)
	http.Get(site)
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("website", site, "online")
	} else {
		fmt.Println("website", site, "offline")
	}
}

//TODO: colocar os sites em um arquivo txt e utilizar leitura de arquivo
