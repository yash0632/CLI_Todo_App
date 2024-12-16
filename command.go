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
	Toggle int
	List   bool
	Edit   string
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add,"add","", "Add a new todo specify title")
	flag.StringVar(&cf.Edit,"edit","","Edit a todo by index & specify a new title, id:new_title")
	flag.IntVar(&cf.Del,"del",-1,"Specify a todo by index to delete")
	flag.IntVar(&cf.Toggle,"toggle",-1,"Specify a todo by index to toggle")
	flag.BoolVar(&cf.List,"list",false,"List all todos")


	flag.Parse();

	return &cf;
}

func (cf *CmdFlags) Execute(todos *Todos){
	switch{
	case cf.List:
		todos.print()
	
	case cf.Add != "":
		todos.add(cf.Add)

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit,":",2)
		if len(parts) != 2{
			fmt.Println("Error,Invalid Edit Format.Please Use id:new_title format")
			os.Exit(1)
		}

		index,err := strconv.Atoi(parts[0]);
		if(err != nil){
			fmt.Printf("Errors:Invalid Index for exit")
			os.Exit(1);
		}

		todos.edit(index,parts[1]);

	case cf.Toggle != -1:
		if(cf.Toggle >= len(*todos)){
			fmt.Printf("Please give Valid Index for Toggle")
		}
		todos.toggle(cf.Toggle)
	
	case cf.Del != -1:
		if(cf.Del >= len(*todos)){
			fmt.Printf("Please give Valid Index for Delete")
		}
		todos.delete(cf.Del)


	default:
		fmt.Println("Invalid Command")
	}

}