package main

import (
	"flag"
	"fmt"
	"os"
)

func handleGreetCommand(args []string) {
	greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
	name := greetCmd.String("name", "World", "Name to greet")
	greetCmd.Parse(args)
	fmt.Printf("Hello, %s!\n", *name)
}

func handleFarewellCommand(args []string) {
	farewellCmd := flag.NewFlagSet("farewell", flag.ExitOnError)
	name := farewellCmd.String("name", "World", "Name to bid farewell to")
	farewellCmd.Parse(args)
	fmt.Printf("Goodbye, %s!\n", *name)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-cli <command> [arguments]")
		fmt.Println("Available commands: greet, farewell")
		os.Exit(1)
	}

	command := os.Args[1]
	remainingArgs := os.Args[2:]

	switch command {
	case "greet":
		handleGreetCommand(remainingArgs)
	case "farewell":
		handleFarewellCommand(remainingArgs)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
