package parkinglot

import "testing"
import "fmt"
import "strings"

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

func TestGetRegNoByColor(t *testing.T) {
	message := testParkinglotMod.CreateParkingLot(capacityTest)
	fmt.Println(message.Message)

	message = testParkinglotMod.ParkCar(&carReg1)

	fmt.Println(message.Message)

	message = testParkinglotMod.GetRegNumberByColor("White")

	fmt.Println(message.Message)

	if message.Message != "KA-01-HH-1234" {
		t.Errorf("Failed to get data")
	}

}

func TestGetSlotByColor(t *testing.T) {
	message := testParkinglotMod.CreateParkingLot(capacityTest)
	fmt.Println(message.Message)

	message = testParkinglotMod.ParkCar(&carReg1)

	fmt.Println(message.Message)

	message = testParkinglotMod.GetSlotByColor("White")

	fmt.Println(message.Message)

	if message.Message != "1" {
		t.Errorf("Failed to get data")
	}

}

func TestGetSlotByCarReg(t *testing.T) {
	message := testParkinglotMod.CreateParkingLot(capacityTest)
	fmt.Println(message.Message)

	message = testParkinglotMod.ParkCar(&carReg1)

	fmt.Println(message.Message)

	message = testParkinglotMod.GetSlotByColor("KA-01-HH-1234")

	fmt.Println(message.Message)

	if message.Message != "1" {
		t.Errorf("Failed to get data")
	}

}

func TestGetStatus(t *testing.T) {
	message := testParkinglotMod.CreateParkingLot(capacityTest)
	fmt.Println(message.Message)

	message = testParkinglotMod.ParkCar(&carReg1)

	fmt.Println(message.Message)

	message = testParkinglotMod.GetStatus()

	fmt.Println(message.Message)

	if !strings.Contains(message.Message, "KA-01-HH-1234") {
		t.Errorf("Check Status doesn't contain correct data")
	}
}

func TestExecuteCommand(t *testing.T) {
	testParkinglotMod.ExecuteCommand(commandCreateTest)

	if testParkinglotMod.Capacity == 0 {
		t.Errorf("Parking Lot is not created with correct capacity")
	}
}

func TestProcessCommand(t *testing.T) {
	testParkinglotMod.ProcessCommands(commandsTest)
	if testParkinglotMod.Capacity == 0 {
		t.Errorf("Parking Lot is not created with correct capacity")
	}
	if testParkinglotMod.ParkingLot.Lots[0] == nil {
		t.Errorf("Car Failed to be added on Parking Lot")
	}
}
