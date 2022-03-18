package main

import "fmt"

func main() {
	var cmd string
	for {
		fmt.Print(">>")
		fmt.Scan(&cmd)
		switch cmd {
		case "help":
			fmt.Println("This is help cmd")
		case "version":
			fmt.Println("Menu v1.0")
		case "quit":
			return
		default:
			fmt.Println("Wrong Command !!!")
		}
	}
}
