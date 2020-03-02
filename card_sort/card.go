package card_sort

import (
	"fmt"
	"strings"
)

type TripType int

const (
	FLIGHT TripType = iota
	BUS
	TRAIN
)

type BoardingCard struct {
	Origin      string
	Destination string
	Trip        TripType

	Seat       string
	VesselInfo string
	Metadata   map[string]string
}

func (bc BoardingCard) String() string {
	var output strings.Builder
	switch bc.Trip {
	case FLIGHT:
		output.WriteString(fmt.Sprintf("From %s, take flight %s to %s. Gate %s, seat %s. Baggage ",
			bc.Origin, bc.VesselInfo, bc.Destination, bc.Metadata["gate"], bc.Seat))
		if counter, ok := bc.Metadata["ticket counter"]; ok {
			output.WriteString(fmt.Sprintf("drop at ticket counter %s.", counter))
		} else {
			output.WriteString(fmt.Sprintf("will be automatically transferred from your last leg."))
		}
	case BUS:
		output.WriteString(fmt.Sprintf("Take the %s from %s to %s.", bc.VesselInfo, bc.Origin, bc.Destination))
		if bc.Seat != "" {
			output.WriteString(fmt.Sprintf(" Sit in seat %s.", bc.Seat))
		} else {
			output.WriteString(" No seat assignment.")
		}
	case TRAIN:
		output.WriteString(fmt.Sprintf("Take train %s from %s to %s.", bc.VesselInfo, bc.Origin, bc.Destination))
		if bc.Seat != "" {
			output.WriteString(fmt.Sprintf(" Sit in seat %s.", bc.Seat))
		} else {
			output.WriteString(" No seat assignment.")
		}
	}
	return output.String()
}
