package main 

func getCommands() map[string] cliCommands{
	return map[string] cliCommands {
		"exit": {
			name: "exit",
			description: "used to exit the cli",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "used to help the users",
			callback: commandHelp,
		},
	}
}