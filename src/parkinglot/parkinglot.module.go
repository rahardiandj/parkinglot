package parkinglot

import (
	"context"
)

// type Mod struct {
// 	Commands []Command
// }

type Mod struct {
	Commands []string
}

func New(ctx context.Context) *Mod {
	parkingLotMod := &Mod{}
	return parkingLotMod
}
