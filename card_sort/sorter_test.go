package card_sort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	t.Parallel()

	// Sample input
	input := []*BoardingCard{
		{
			Origin:      "Barcelona",
			Destination: "Gerona Airport",
			Trip:        BUS,
			Seat:        "",
			VesselInfo:  "airport bus",
			Metadata:    map[string]string{},
		},
		{
			Origin:      "Stockholm",
			Destination: "New York JFK",
			Trip:        FLIGHT,
			Seat:        "7B",
			VesselInfo:  "SK22",
			Metadata: map[string]string{
				"gate": "22",
			},
		},
		{
			Origin:      "Madrid",
			Destination: "Barcelona",
			Trip:        TRAIN,
			Seat:        "45B",
			VesselInfo:  "78A",
			Metadata:    map[string]string{},
		},
		{
			Origin:      "Gerona Airport",
			Destination: "Stockholm",
			Trip:        FLIGHT,
			Seat:        "3A",
			VesselInfo:  "SK455",
			Metadata: map[string]string{
				"gate":           "45B",
				"ticket counter": "344",
			},
		},
	}

	// Define expectations
	want := []*BoardingCard{
		{
			Origin:      "Madrid",
			Destination: "Barcelona",
			Trip:        TRAIN,
			Seat:        "45B",
			VesselInfo:  "78A",
			Metadata:    map[string]string{},
		},
		{
			Origin:      "Barcelona",
			Destination: "Gerona Airport",
			Trip:        BUS,
			Seat:        "",
			VesselInfo:  "airport bus",
			Metadata:    map[string]string{},
		},
		{
			Origin:      "Gerona Airport",
			Destination: "Stockholm",
			Trip:        FLIGHT,
			Seat:        "3A",
			VesselInfo:  "SK455",
			Metadata: map[string]string{
				"gate":           "45B",
				"ticket counter": "344",
			},
		},
		{
			Origin:      "Stockholm",
			Destination: "New York JFK",
			Trip:        FLIGHT,
			Seat:        "7B",
			VesselInfo:  "SK22",
			Metadata: map[string]string{
				"gate": "22",
			},
		},
	}
	wantErr := error(nil)

	s := New()
	got, gotErr := s.Sort(input)
	if gotErr != wantErr {
		t.Errorf("An unexpected error was returned. Got: %v, Want: %v", gotErr, wantErr)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("The boarding cards were not sorted correctly. Got: %+v, Want: %+v", got, want)
	}
}

func TestSortHasCycle(t *testing.T) {
	t.Parallel()

	// Sample input
	input := []*BoardingCard{
		{
			Origin:      "Stockholm",
			Destination: "New York JFK",
			Trip:        FLIGHT,
			Seat:        "7B",
			VesselInfo:  "SK22",
			Metadata: map[string]string{
				"gate": "22",
			},
		},
		{
			Origin:      "Madrid",
			Destination: "Barcelona",
			Trip:        TRAIN,
			Seat:        "45B",
			VesselInfo:  "78A",
			Metadata:    map[string]string{},
		},
		{
			Origin:      "Gerona Airport",
			Destination: "Stockholm",
			Trip:        FLIGHT,
			Seat:        "3A",
			VesselInfo:  "SK455",
			Metadata: map[string]string{
				"gate":           "45B",
				"ticket counter": "344",
			},
		},
		{
			Origin:      "Barcelona",
			Destination: "Madrid",
			Trip:        FLIGHT,
			Seat:        "23A",
			VesselInfo:  "SK23",
			Metadata: map[string]string{
				"gate":           "12",
				"ticket counter": "A",
			},
		},
	}

	// Define expectations
	want := []*BoardingCard(nil)
	wantErr := ErrHasCycle

	s := New()
	got, gotErr := s.Sort(input)
	if gotErr != wantErr {
		t.Errorf("An unexpected error was returned. Got: %v, Want: %v", gotErr, wantErr)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("The boarding cards were not sorted correctly. Got: %+v, Want: %+v", got, want)
	}
}

func TestSortSmallInput(t *testing.T) {
	t.Parallel()

	// Sample input
	input := []*BoardingCard{
		{
			Origin:      "Barcelona",
			Destination: "Gerona Airport",
			Trip:        BUS,
			Seat:        "",
			VesselInfo:  "airport bus",
			Metadata:    map[string]string{},
		},
	}

	// Define expectations
	want := []*BoardingCard{
		{
			Origin:      "Barcelona",
			Destination: "Gerona Airport",
			Trip:        BUS,
			Seat:        "",
			VesselInfo:  "airport bus",
			Metadata:    map[string]string{},
		},
	}
	wantErr := error(nil)

	s := New()
	got, gotErr := s.Sort(input)
	if gotErr != wantErr {
		t.Errorf("An unexpected error was returned. Got: %v, Want: %v", gotErr, wantErr)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("The boarding cards were not sorted correctly. Got: %+v, Want: %+v", got, want)
	}
}

func TestSortNilInput(t *testing.T) {
	t.Parallel()

	// Sample input
	input := []*BoardingCard(nil)

	// Define expectations
	want := []*BoardingCard(nil)
	wantErr := error(nil)

	s := New()
	got, gotErr := s.Sort(input)
	if gotErr != wantErr {
		t.Errorf("An unexpected error was returned. Got: %v, Want: %v", gotErr, wantErr)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("The boarding cards were not sorted correctly. Got: %+v, Want: %+v", got, want)
	}
}
