package parkinglot

import (
	"fmt"
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
			result = append(result, string(i))
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

func (mod *Mod) GetSlotByRegNo(regNo string) Message {
	var result []string

	for i, lot := range mod.ParkingLot.Lots {
		if lot.Number == regNo {
			result = append(result, string(i))
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
	result := fmt.Sprintf("Slot No.	Registrtaion No		Colour\n")
	for i, lot := range mod.ParkingLot.Lots {
		if lot != nil {
			result = fmt.Sprintf("%v%v		%v		%v\n", result, i+1, lot.Number, lot.Color)
		}
	}
	message := Message{
		Message: result,
	}
	return message
}
