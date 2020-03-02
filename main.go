package main

import (
	"fmt"

	"github.com/kylechadha/card-sort-challenge/card_sort"
)

func main() {
	sampleInput := []card_sort.BoardingCard{
		{
			Origin:      "Barcelona",
			Destination: "Gerona Airport",
			Seat:        "",
			Gate:        "",
			Vessel:      "airport bus",
			MetaData: card_sort.TripMetadata{
				Trip: card_sort.BUS,
			},
		},
		{
			Origin:      "Stockholm",
			Destination: "New York JFK",
			Seat:        "7B",
			Gate:        "22",
			Vessel:      "SK22",
			MetaData: card_sort.TripMetadata{
				Trip: card_sort.FLIGHT,
			},
		},
		{
			Origin:      "Madrid",
			Destination: "Barcelona",
			Seat:        "45B",
			Gate:        "",
			Vessel:      "78A",
			MetaData: card_sort.TripMetadata{
				Trip: card_sort.TRAIN,
			},
		},
		{
			Origin:        "Gerona Airport",
			Destination:   "Stockholm",
			Seat:          "3A",
			Gate:          "45B",
			Vessel:        "SK455",
			TicketCounter: "344",
			MetaData: card_sort.TripMetadata{
				Trip: card_sort.FLIGHT,
			},
		},
	}

	s := card_sort.New()
	sortMap, origin := s.SortCards(sampleInput)
	r := card_sort.NewCardReader()
	output := r.ReadCards(sortMap, origin)
	fmt.Println(output)
}
