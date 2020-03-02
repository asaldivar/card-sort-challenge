package card_sort

// CardSorter accepts a slice of cards and sorts them
// with O(n) time and O(n) space.
type CardSorter interface {
	SortCards(cards []BoardingCard) (sortMap map[string]BoardingCard, origin string)
}

type sorter struct {
	sortMap map[string]BoardingCard
}

func New() CardSorter {
	return sorter{sortMap: make(map[string]BoardingCard)}
}

// Pseudocode
// - We want to find the source node in a directed graph
// - The algorithm for this is topological sort
// - The last vertex with no edges coming out gives us the sink
// - So in order to find the source, we just need to reverse the graph
func (s sorter) SortCards(cards []BoardingCard) (sortMap map[string]BoardingCard, origin string) {
	// Linear time and space.
	cardMap := make(map[string]BoardingCard)
	for _, card := range cards {
		cardMap[card.Origin] = card
	}

	// Linear time and space.
	reverseGraph := make(map[string]string)
	for _, card := range cards {
		s.sortMap[card.Origin] = cardMap[card.Destination]
		reverseGraph[card.Destination] = card.Origin
	}

	for dest, orig := range reverseGraph {

	}
}
