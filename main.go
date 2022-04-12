package main

import "fmt"

func main() {
	cmdList := initalCmdList()
	for {
		fmt.Print(">>")
		var cmdName string
		fmt.Scan(&cmdName)
		cmdList.exec(cmdName)
	}
}

func initalCmdList() *CmdList {
	cmdList := getCmdList()
	cmdList.insert("help", "show manual", &Help{"help", cmdList}, true)
	cmdList.insert("version", "show version", &Version{"Version"}, true)
	cmdList.insert("quit", "quit this menu", &Quit{"Quit"}, true)
	return cmdList
}
