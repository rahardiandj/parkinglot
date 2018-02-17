package main

import (
	"context"
	"fmt"

	"github.com/rahardiandj/parkingLot/src/parkinglot"
)

func main() {

	ctx := context.Background()
	parkinglotMod := parkinglot.New(ctx)
	parkinglotMod.Init(ctx)

	car := parkinglot.CarRegistration{
		Number: "K-02323-KK",
		Color:  "White",
	}
	message := parkinglotMod.CreateParkingLot(3)
	fmt.Printf("%s\n", message.Message)
	message = parkinglotMod.ParkCar(&car)
	fmt.Printf("%s\n", message.Message)

	message = parkinglotMod.ParkCar(&car)
	fmt.Printf("%s\n", message.Message)

	message = parkinglotMod.ParkCar(&car)
	fmt.Printf("%s\n", message.Message)

	message = parkinglotMod.LeaveLot(2)
	fmt.Printf("%s\n", message.Message)

	message = parkinglotMod.GetStatus()
	fmt.Printf("%s\n", message.Message)
}
