package parkinglot

import (
	"context"
	"fmt"
)

func (mod *Mod) Init(context.Context) {
	mod.Commands = InitCommand()

	if mod.Commands != nil {
		for i, command := range mod.Commands {
			fmt.Printf("%v. %v %v\n", i+1, command.CommandStr, command.Argument1)
		}
	}
}

func InitCommand() []Command {
	var commands []Command

	create := Command{
		CommandStr: "create_parking_lot",
		Argument1:  "[Capacity]",
	}

	park := Command{
		CommandStr: "park",
		Argument1:  "[Register No] [Color]",
	}

	leave := Command{
		CommandStr: "leave",
		Argument1:  "[Slot Number]",
	}

	status := Command{
		CommandStr: "status",
	}

	checkRegNoByColor := Command{
		CommandStr: "registration_numbers_for_cars_with_colour",
		Argument1:  "[Colour]",
	}

	checkSlotByColor := Command{
		CommandStr: "slot_numbers_for_cars_with_colour",
		Argument1:  "[Register No]",
	}

	checkSlotByRegNo := Command{
		CommandStr: "slot_number_for_registration_number",
		Argument1:  "[Register No]",
	}

	commands = append(commands, create, park, leave, status, checkRegNoByColor, checkSlotByColor, checkSlotByRegNo)
	return commands
}

// func InitMessage() []Message {
// 	var messages []Message
// 	created := Message{
// 		Message: `Created a parking lot with %v slots`,
// 	}
// 	allocated := Message{
// 		Message: "Allocated slot number: %v",
// 	}
// 	slotfree := Message{
// 		Message: "Slot number %v is free",
// 	}
// 	full := Message{
// 		Message: "Sorry, parking lot is full",
// 	}
// 	notfound := Message{
// 		Message: "Not Found",
// 	}
// 	messages = append(messages, created, allocated, slotfree, full, notfound)

// 	return messages
// }
