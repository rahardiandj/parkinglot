package parkinglot

import (
	"context"
)

func (mod *Mod) Init(context.Context) {
	mod.Commands = InitCommand()
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

	exit := Command{
		CommandStr: "exit",
	}

	commands = append(commands, create, park, leave, status, checkRegNoByColor, checkSlotByColor, checkSlotByRegNo, exit)
	return commands
}
