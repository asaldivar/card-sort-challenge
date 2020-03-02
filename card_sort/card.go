package card_sort

const (
	FLIGHT TripType = iota
	BUS
	TRAIN
)

type BoardingCard struct {
	Origin      string
	Destination string

	Seat          string
	Gate          string
	Vessel        string
	TicketCounter string
	MetaData      TripMetadata
}

// TODO: Replace TripMetadata struct with MetaProvider
// interface that can hold any metadata for trips.
type MetaProvider interface {
	GetAll()
	Get(name string)
	Trip() TripType
}

type TripMetadata struct {
	Trip TripType
	// other metadata can be added here
}

type TripType int
