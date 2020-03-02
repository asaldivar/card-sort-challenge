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
			Trip:        card_sort.BUS,
			Seat:        "",
			VesselInfo:  "airport bus",
			Metadata:    map[string]string{},
		},
		{
			Origin:      "Stockholm",
			Destination: "New York JFK",
			Trip:        card_sort.FLIGHT,
			Seat:        "7B",
			VesselInfo:  "SK22",
			Metadata: map[string]string{
				"gate": "22",
			},
		},
		{
			Origin:      "Madrid",
			Destination: "Barcelona",
			Trip:        card_sort.TRAIN,
			Seat:        "45B",
			VesselInfo:  "78A",
			Metadata:    map[string]string{},
		},
		{
			Origin:      "Gerona Airport",
			Destination: "Stockholm",
			Trip:        card_sort.FLIGHT,
			Seat:        "3A",
			VesselInfo:  "SK455",
			Metadata: map[string]string{
				"gate":           "45B",
				"ticket counter": "344",
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
