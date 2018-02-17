package parkinglot

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (mod *Mod) ProcessFile() {
	arg := os.Args[1]
	commands := mod.FileReader(arg)
	if commands == nil {
		fmt.Printf("File text is empty.\n")
	}
	mod.ProcessCommands(commands)
}

func (mod *Mod) ProcessCommand() {
	if mod.Commands != nil {
		fmt.Printf("Provided Command: \n")
		for i, command := range mod.Commands {
			fmt.Printf("%v. %v %v\n", i+1, command.CommandStr, command.Argument1)
		}
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("\nInput (type exit to exit): \n")
	for scanner.Scan() {

		var command Command

		line := scanner.Text()
		if line == "exit" {
			os.Exit(0)
		}

		strParam := strings.Fields(line)

		command.CommandStr = strParam[0]
		if len(strParam) > 2 {
			command.Argument2 = strParam[2]
		}
		if len(strParam) > 1 {
			command.Argument1 = strParam[1]
		}
		fmt.Printf("\nOutput : \n")
		mod.ExecuteCommand(command)

		fmt.Printf("\nInput (type exit to exit): \n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
func (mod *Mod) FileReader(filePath string) []Command {
	file, err := os.Open(filePath)
	var commands = []Command{}
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var command Command
		line := fmt.Sprintf(scanner.Text())
		strParam := strings.Fields(line)

		command.CommandStr = strParam[0]
		if len(strParam) > 2 {
			command.Argument2 = strParam[2]
		}
		if len(strParam) > 1 {
			command.Argument1 = strParam[1]
		}
		commands = append(commands, command)
	}
	return commands
}

func (m *Message) Formatted(Argument1 string) {
	m.Message = fmt.Sprintf(m.Message, Argument1)
}
