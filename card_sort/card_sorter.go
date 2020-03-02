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
// - The highest post order value is the source

// - For the sake of time:
// - I think we can get around this by reversing the graph
//   and just seeing if there's any origins that aren't in
//   the reverse graph (meaning they have no edges coming out)
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

	// TODO: Do we need to handle cycles? Are there islands? Ie: Can we always
	// assume the input is a valid connected DAG?
	for orig, destCard := range cardMap {
		if _, ok := reverseGraph[orig]; !ok {
			origin = orig
		}
	}

	return sortMap, origin
}
