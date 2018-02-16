package parkinglot

import (
	"context"
	"os"
	"testing"
)

var capacityTest int
var testParkinglotMod *Mod
var carReg1 CarRegistration

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
}
