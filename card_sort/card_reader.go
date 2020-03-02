package card_sort

import "fmt"

// CardReader accepts a map of origin to destinations
// and an origin and prints an itinerary to standard out.
//
// TODO: Change this interface to conform to io.Reader pattern.
type CardReader interface {
	ReadCards(cardMap map[string]BoardingCard, origin BoardingCard) []BoardingCard
}

type reader struct {
}

func NewCardReader() CardReader {
	return reader{}
}

func (r reader) ReadCards(cardMap map[string]BoardingCard, originCard BoardingCard) []BoardingCard {
	output := []BoardingCard{originCard}
	next := originCard.Origin
	fmt.Println(len(cardMap))
	for i := 0; i < len(cardMap)-1; i++ {
		output = append(output, cardMap[next])
		next = cardMap[next].Origin
	}
	return output
}
