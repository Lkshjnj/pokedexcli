package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommands struct {
	name string
	description string
	callback func() error
}


func main() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the REPL ! type 'exit' to quit.")
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan(){
			fmt.Println("errors while taking the input !")
			break
		}else {
			inputText := scanner.Text()
			cleansedInput := cleanInput(inputText)
			if _, ok := commands[cleansedInput[0]];!ok{
				fmt.Println("Unkown command, please try again")
				continue
			} else{
				commands[cleansedInput[0]].callback()
			}
		}
	}

}
