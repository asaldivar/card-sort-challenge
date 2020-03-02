package card_sort

const (
	FLIGHT TripType = iota
	BUS
	TRAIN
)

type BoardingCard struct {
	Origin        string
	Destination   string
	Seat          string
	Gate          string
	Vessel        string
	TicketCounter string
	MetaData      TripMetadata
}

type MetaProvider interface {
	GetAll()
	Get(name string)
}

type TripMetadata struct {
	Trip TripType
	// other metadata can be added here
}

type TripType int
