package main

import (
	"fmt"
	"os"
)

type Help struct {
	cmdName string
	cmdList *CmdList
}

func (this *Help) exec() {
	this.cmdList.printCmd()
}

type Quit struct {
	cmdName string
}

func (this *Quit) exec() {
	os.Exit(0)
}

type Version struct {
	cmdName string
}

func (this *Version) exec() {
	fmt.Println("Menu v2.0")
}
