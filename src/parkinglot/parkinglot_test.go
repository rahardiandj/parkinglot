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
