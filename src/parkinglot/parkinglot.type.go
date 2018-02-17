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
	Message string
	Result  interface{}
}

const (
	STATUS_AVAILABLE  = "AVAILABLE"
	STATUS_FULL       = "FULL"
	STATUS_NOT_FOUND  = "NOTFOUND"
	MESSAGE_FULL      = "Sorry, parking lot is full"
	MESSAGE_NOT_FOUND = "Not Found"
)

type Command struct {
	CommandStr string
	Argument   string
}
