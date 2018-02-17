package main

import (
	"context"
	"os"

	"github.com/rahardiandj/parkingLot/src/parkinglot"
)

func main() {

	ctx := context.Background()
	parkinglotMod := parkinglot.New(ctx)
	parkinglotMod.Init(ctx)

	if len(os.Args) > 1 {
		parkinglotMod.ProcessFile()
	} else {
		parkinglotMod.ProcessCommand()
	}

}
