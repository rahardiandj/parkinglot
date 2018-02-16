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
	message := parkinglotMod.CreateParkingLot(1)
	fmt.Printf("%s\n", message.Message)
	message = parkinglotMod.ParkCar(&car)
	fmt.Printf("%s\n", message.Message)

	message = parkinglotMod.ParkCar(&car)
	fmt.Printf("%s\n", message.Message)

}
