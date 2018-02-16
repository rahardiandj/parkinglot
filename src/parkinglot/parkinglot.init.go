package parkinglot

func (mod *Mod) Init() {
	//commands := InitCommand()
	//mod.Commands = commands
}

func InitCommand() []Command {
	var commands []Command

	create := Command{
		CommandStr: "create_parking_lot",
	}

	park := Command{
		CommandStr: "park",
	}

	leave := Command{
		CommandStr: "leave",
	}

	status := Command{
		CommandStr: "status",
	}

	checkRegNoByColor := Command{
		CommandStr: "registration_numbers_for_cars_with_colour",
	}

	checkSlotByColor := Command{
		CommandStr: "slot_numbers_for_cars_with_colour",
	}

	checkSlotByRegNo := Command{
		CommandStr: "slot_number_for_registration_number",
	}
	commands = append(commands, create, park, leave, status, checkRegNoByColor, checkSlotByColor, checkSlotByRegNo)
	return commands
}
