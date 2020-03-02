package card_sort

import "fmt"

// CardReader accepts a map of origin to destinations
// and an origin and prints an itinerary to standard out.
//
// TODO: Change this interface to conform to io.Reader pattern.
type CardReader interface {
	ReadCards(cardMap map[string]BoardingCard, origin string)
}

type reader struct {
}

func NewCardReader() CardReader {
	return reader{}
}

func (r reader) ReadCards(cardMap map[string]BoardingCard, origin string) {
	next := origin
	for i := 0; i < len(cardMap); i++ {
		fmt.Println(cardMap[next])
		next = cardMap[next].Destination
	}
}
