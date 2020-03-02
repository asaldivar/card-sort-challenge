package main

import (
	"fmt"
	"os"

	"github.com/kylechadha/card-sort-challenge/card_sort"
)

func main() {
	input := []*card_sort.BoardingCard{
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
	output, err := s.Sort(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Sorting could not be completed: %v\n", err)
		os.Exit(1)
	}

	for _, card := range output {
		fmt.Println(*card)
	}
}
