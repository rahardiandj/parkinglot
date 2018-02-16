package parkinglot

import (
	"fmt"
)

func (mod *Mod) CreateParkingLot(capacity int) Message {
	mod.ParkingLot.Capacity = capacity
	mod.Message[0].Formatted()
	return mod.Message[0]
}

func (m *Message) Formatted() {
	m.Message = fmt.Sprintf(m.Message, m.Argument)
}
