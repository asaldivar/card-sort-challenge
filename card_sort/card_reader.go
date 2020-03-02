package card_sort

// CardReader accepts a map of origin to destinations
// and an origin and prints an itinerary to standard out.
//
// TODO: Change this interface to conform to io.Reader pattern.
type CardReader interface {
	ReadCards(cardMap map[string]BoardingCard, origin string) []BoardingCard
}

type reader struct {
}
