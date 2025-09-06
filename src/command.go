package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add new task into the Todo list - 'title'")
	flag.StringVar(&cf.Edit, "edit", "", "Edit task in the Todo list at specified id - 'id:new_title'")
	flag.IntVar(&cf.Del, "del", -1, "Delete task at specified id from the Todo list - id")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a task at specified id from the Todo list - id")
	flag.BoolVar(&cf.List, "list", false, "List all tasks from the Todo list")
	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()

	case cf.Add != "":
		todos.add(cf.Add)
		todos.print()

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit. Please use 'id:new_title'")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: invalid id for edit")
			os.Exit(1)
		}
		todos.edit(index, parts[1])
		todos.print()

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
		todos.print()

	case cf.Del != -1:
		todos.delete(cf.Del)
		todos.print()

	default:
		fmt.Println("Invalid command")
	}
}
