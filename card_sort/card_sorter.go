package card_sort

import (
	"errors"
)

var ErrHasCycle = errors.New("the itinerary contains a cycle")

// CardSorter accepts a slice of *BoardingCard and sorts them based on
// origin and destination information.
type CardSorter interface {
	Sort(cards []*BoardingCard) ([]*BoardingCard, error)
}

// sorter sorts a slice of *BoardingCard using topological sort, in
// O(|V|+|E|) time and O(|V|+|E|) space, where |V| is the number of
// nodes and |E| is the number of edges. The sort is NOT conducted
// in place.
//
// If the slice of *BoardingCard passed in does not represent a valid
// DAG, ie: it contains a cycle, an error is returned.
type sorter struct{}

func New() CardSorter {
	return sorter{}
}

func (s sorter) Sort(cards []*BoardingCard) ([]*BoardingCard, error) {
	cardMap := make(map[string]*BoardingCard)
	for _, card := range cards {
		cardMap[card.Origin] = card
	}

	adjList := make(map[string][]*BoardingCard)
	for _, card := range cards {
		if dest, ok := cardMap[card.Destination]; ok {
			adjList[card.Origin] = append(adjList[card.Origin], dest)
		}
	}

	var output []*BoardingCard
	var topoSort func(node string, adjList map[string][]*BoardingCard, visited map[string]int) error
	topoSort = func(node string, adjList map[string][]*BoardingCard, visited map[string]int) error {
		if visited[node] == -1 {
			// We hit a node who's neighbors are still being explored. There's a cycle.
			return ErrHasCycle
		}
		if visited[node] == 1 {
			// We've already explored this node.
			return nil
		}

		visited[node] = -1
		for _, dest := range adjList[node] {
			err := topoSort(dest.Origin, adjList, visited)
			if err != nil {
				return err
			}
		}
		visited[node] = 1

		output = append(output, cardMap[node])
		return nil
	}

	// Visited tracks the exploration state of nodes.
	//   0 for nodes that have not been visited
	//   1 for nodes that have been visited and fully explored
	//  -1 for nodes that have been visited and are still being explored
	visited := make(map[string]int)
	for _, card := range cards {
		err := topoSort(card.Origin, adjList, visited)
		if err != nil {
			return nil, err
		}
	}

	// Reverse the output.
	start, end := 0, len(output)-1
	for start < end {
		output[start], output[end] = output[end], output[start]
		start++
		end--
	}
	return output, nil
}
