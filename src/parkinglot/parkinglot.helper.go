package parkinglot

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func (m *Message) Formatted(argument string) {
	m.Message = fmt.Sprintf(m.Message, argument)
}
