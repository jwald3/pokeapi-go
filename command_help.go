package main

import "fmt"

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the pokedex help menu.")
	fmt.Println("Here are the available commands:")

	availableCommands := getCommands()

	for _, v := range availableCommands {
		fmt.Printf(" - %v\t\t%v\n", v.name, v.description)
	}
	fmt.Println()

	return nil
}
