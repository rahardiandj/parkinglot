package parkinglot

import (
	"context"
)

type Mod struct {
	Commands []Command
	Message  []Message
	ParkingLot
}

func New(ctx context.Context) *Mod {
	parkinglotMod := &Mod{}
	return parkinglotMod
}
