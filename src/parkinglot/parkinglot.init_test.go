package parkinglot

import (
	"context"
	"os"
	"testing"
)

var capacityTest int
var testParkinglotMod *Mod
var carReg1 CarRegistration
var commandCreateTest, commandParkTest Command
var commandsTest []Command

func TestMain(m *testing.M) {

	initData()
	ctx := context.Background()
	testParkinglotMod = New(ctx)
	testParkinglotMod.Init(context.Background())

	os.Exit(m.Run())
}

func initData() {
	capacityTest = 3
	carReg1.Number = "KA-01-HH-1234"
	carReg1.Color = "White"

	commandCreateTest.CommandStr = "create_parking_lot"
	commandCreateTest.Argument1 = "6"

	commandParkTest.CommandStr = "park"
	commandParkTest.Argument1 = carReg1.Number
	commandCreateTest.Argument2 = carReg1.Color

	commandsTest = append(commandsTest, commandCreateTest, commandParkTest)
}
