package parkinglot

type CarRegistration struct {
	Number string
	Color  string
}

type ParkingLot struct {
	Lots     []*CarRegistration
	Capacity int
	Status   string
}

type Message struct {
	Message  string
	Argument string
}

const (
	STATUS_AVAILABLE = "AVAILABLE"
	STATUS_FULL      = "FULL"
	STATUS_NOT_FOUND = "NOTFOUND"
)

type Command struct {
	CommandStr string
	Argument   string
}
