package parkinglot

import (
	"context"
	"os"
	"testing"
)

var capacityTest int
var testParkinglotMod *Mod

func TestMain(m *testing.M) {

	initData()
	ctx := context.Background()
	testParkinglotMod = New(ctx)
	testParkinglotMod.Init(context.Background())

	os.Exit(m.Run())
}

func initData() {
	capacityTest = 3
}
