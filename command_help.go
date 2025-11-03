package main 

import (
	"fmt"
)

func commandHelp() error{
	fmt.Printf("[+] The available commands are as follows:\n\n")
	fmt.Println("--------------------------------------------------------------------------------")
	commands := getCommands()
	for _, item := range commands{
		fmt.Printf("command name: %s\n", item.name)
		fmt.Printf("Description: %s\n", item.description)
		fmt.Println("--------------------------------------------------------------------------------")
	}
	return nil
}