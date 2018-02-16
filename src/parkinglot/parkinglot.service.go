package parkinglot

import "strconv"
import "fmt"

func (mod *Mod) CreateParkingLot(capacity int) Message {
	lots := make([]*CarRegistration, capacity)
	mod.ParkingLot.Capacity = capacity
	mod.ParkingLot.Lots = lots
	mod.Status = STATUS_AVAILABLE
	mod.Message[0].Formatted(string(capacity))
	return mod.Message[0]
}

func (mod *Mod) ParkCar(carReg *CarRegistration) Message {
	fmt.Println(mod.Status)
	if mod.Status == STATUS_FULL {
		return mod.Message[3]
	}
	for i, lot := range mod.ParkingLot.Lots {
		if lot == nil {
			mod.ParkingLot.Lots[i] = carReg
			mod.Message[1].Formatted(strconv.Itoa(i))
			return mod.Message[1]
		}
	}
	mod.ParkingLot.Status = STATUS_FULL
	return mod.Message[3]
}

func (mod *Mod) LeaveLot(slotNo int) Message {

	if mod.ParkingLot.Lots[slotNo-1] == nil {
		return mod.Message[4]
	}
	mod.ParkingLot.Lots[slotNo-1] = nil
	if mod.ParkingLot.Status == STATUS_FULL {
		mod.ParkingLot.Status = STATUS_AVAILABLE
	}
	mod.Message[2].Formatted(string(slotNo))
	return mod.Message[2]
}
