package parkinglot

import "testing"
import "fmt"

func TestCreateParkingLot(t *testing.T) {

	message := testParkinglotMod.CreateParkingLot(capacityTest)
	fmt.Println(message.Message)

	if testParkinglotMod.Capacity != capacityTest {
		t.Errorf("Parking Lot is not created with correct capacity")
	}

}

func TestParkCar(t *testing.T) {
	message := testParkinglotMod.CreateParkingLot(capacityTest)
	fmt.Println(message.Message)

	message = testParkinglotMod.ParkCar(&carReg1)

	fmt.Println(message.Message)

	if testParkinglotMod.ParkingLot.Lots[0] == nil {
		t.Errorf("Car Failed to be added on Parking Lot")
	}

}

func TestLeaveLot(t *testing.T) {
	message := testParkinglotMod.CreateParkingLot(capacityTest)
	fmt.Println(message.Message)

	message = testParkinglotMod.ParkCar(&carReg1)

	fmt.Println(message.Message)

	message = testParkinglotMod.LeaveLot(1)

	fmt.Println(message.Message)

	if testParkinglotMod.ParkingLot.Lots[0] != nil {
		t.Errorf("Car Failed to be added on Parking Lot")
	}

}
