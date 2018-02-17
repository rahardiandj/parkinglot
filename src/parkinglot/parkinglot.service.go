package parkinglot

import (
	"fmt"
	"strconv"
	"strings"
)

func (mod *Mod) CreateParkingLot(capacity int) Message {
	lots := make([]*CarRegistration, capacity)

	mod.ParkingLot.Capacity = capacity
	mod.ParkingLot.Lots = lots
	mod.Status = STATUS_AVAILABLE

	message := Message{
		Message: fmt.Sprintf(`Created a parking lot with %v slots`, capacity),
		Result:  capacity,
	}
	return message
}

func (mod *Mod) ParkCar(carReg *CarRegistration) Message {
	message := Message{
		Message: MESSAGE_FULL,
	}

	if mod.Status == STATUS_FULL {
		return message
	}
	for i, lot := range mod.ParkingLot.Lots {
		if lot == nil {
			mod.ParkingLot.Lots[i] = carReg

			message := Message{
				Message: fmt.Sprintf("Allocated slot number: %v", i+1),
				Result:  carReg,
			}

			return message
		}
	}
	mod.ParkingLot.Status = STATUS_FULL
	return message
}

func (mod *Mod) LeaveLot(slotNo int) Message {

	message := Message{
		Message: MESSAGE_NOT_FOUND,
	}

	if mod.ParkingLot.Lots[slotNo-1] == nil {
		return message
	}

	mod.ParkingLot.Lots[slotNo-1] = nil
	if mod.ParkingLot.Status == STATUS_FULL {
		mod.ParkingLot.Status = STATUS_AVAILABLE
	}

	message.Message = fmt.Sprintf("Slot number %v is free", slotNo)

	return message
}

func (mod *Mod) GetRegNumberByColor(color string) Message {
	var result []string

	for _, lot := range mod.ParkingLot.Lots {
		if lot.Color == color {
			result = append(result, lot.Number)
		}
	}
	if len(result) == 0 {
		message := Message{
			Message: MESSAGE_NOT_FOUND,
		}
		return message
	}
	message := Message{
		Message: fmt.Sprintf(strings.Join(result[:], ",")),
	}
	return message
}

func (mod *Mod) GetSlotByColor(color string) Message {
	var result []string

	for i, lot := range mod.ParkingLot.Lots {
		if lot.Color == color {
			result = append(result, strconv.Itoa(i+1))
		}
	}
	if len(result) == 0 {
		message := Message{
			Message: MESSAGE_NOT_FOUND,
		}
		return message
	}
	message := Message{
		Message: fmt.Sprintf(strings.Join(result[:], ", ")),
	}
	return message
}

func (mod *Mod) GetSlotByRegNo(regNo string) Message {
	var result []string

	for i, lot := range mod.ParkingLot.Lots {
		if lot.Number == regNo {
			result = append(result, strconv.Itoa(i+1))
		}
	}
	if len(result) == 0 {
		message := Message{
			Message: MESSAGE_NOT_FOUND,
		}
		return message
	}
	message := Message{
		Message: fmt.Sprintf(strings.Join(result[:], ",")),
	}
	return message
}

func (mod *Mod) GetStatus() Message {
	result := fmt.Sprintf("Slot No.	Registration No		Colour")
	for i, lot := range mod.ParkingLot.Lots {
		if lot != nil {
			result = fmt.Sprintf("%v\n%v		%v		%v", result, i+1, lot.Number, lot.Color)
		}
	}
	message := Message{
		Message: result,
	}
	return message
}

func (mod *Mod) ExecuteCommand(command Command) {

	switch command.CommandStr {
	case mod.Commands[0].CommandStr:
		{
			arg, err := strconv.Atoi(command.Argument1)
			if err != nil {
				fmt.Println(err)
			}
			message := mod.CreateParkingLot(arg)
			fmt.Println(message.Message)
		}

	case mod.Commands[1].CommandStr:
		{
			carReg := &CarRegistration{
				Number: command.Argument1,
				Color:  command.Argument2,
			}
			message := mod.ParkCar(carReg)
			fmt.Println(message.Message)
		}
	case mod.Commands[2].CommandStr:
		{
			arg, err := strconv.Atoi(command.Argument1)
			if err != nil {
				fmt.Println(err)
			}
			message := mod.LeaveLot(arg)
			fmt.Println(message.Message)
		}
	case mod.Commands[3].CommandStr:
		{
			message := mod.GetStatus()
			fmt.Println(message.Message)
		}
	case mod.Commands[4].CommandStr:
		{
			message := mod.GetRegNumberByColor(command.Argument1)
			fmt.Println(message.Message)
		}

	case mod.Commands[5].CommandStr:
		{
			message := mod.GetSlotByColor(command.Argument1)
			fmt.Println(message.Message)
		}

	case mod.Commands[6].CommandStr:
		{
			message := mod.GetSlotByRegNo(command.Argument1)
			fmt.Println(message.Message)
		}

	}

}

func (mod *Mod) ProcessCommands(commands []Command) {
	for _, command := range commands {
		mod.ExecuteCommand(command)
	}
}
