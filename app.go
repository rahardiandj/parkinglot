package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rahardiandj/parkingLot/src/parkinglot"
)

func main() {

	ctx := context.Background()
	parkinglotMod := parkinglot.New(ctx)
	parkinglotMod.Init(ctx)

	if len(os.Args) > 1 {
		arg := os.Args[1]
		commands := parkinglotMod.FileReader(arg)
		if commands == nil {
			fmt.Printf("File text is empty.\n")
		}
		parkinglotMod.ProcessCommands(commands)
	}

	// car := parkinglot.CarRegistration{
	// 	Number: "K-02323-KK",
	// 	Color:  "White",
	// }
	// message := parkinglotMod.CreateParkingLot(3)
	// fmt.Printf("%s\n", message.Message)
	// message = parkinglotMod.ParkCar(&car)
	// fmt.Printf("%s\n", message.Message)

	// message = parkinglotMod.ParkCar(&car)
	// fmt.Printf("%s\n", message.Message)

	// message = parkinglotMod.ParkCar(&car)
	// fmt.Printf("%s\n", message.Message)

	// message = parkinglotMod.LeaveLot(2)
	// fmt.Printf("%s\n", message.Message)

	// message = parkinglotMod.GetStatus()
	// fmt.Printf("%s\n", message.Message)
}
