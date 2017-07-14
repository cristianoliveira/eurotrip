package api

import (
	"github.com/cristianoliveira/eurotrip/common"
	"testing"
)

func TestFindingAnUnion(t *testing.T) {
	route := NewRoute([]string{"3", "1", "6", "5"})

	t.Run("It has connection between 1 and 2", func(tt *testing.T) {
		if !route.Connected("3", "6") {
			tt.Errorf("Union between 3 and 6 not found")
		}
	})

	t.Run("It has connection between 5 and 3", func(tt *testing.T) {
		if !route.Connected("5", "3") {
			tt.Errorf("Union between 10 and 2 not found")
		}
	})

	t.Run("It has no connection between 10 and 50", func(tt *testing.T) {
		if route.Connected("1", "50") {
			tt.Errorf("Union between 1, 50 found")
		}
	})
}

func TestLoadItinerary(t *testing.T) {
	settings := common.Settings()
	itineraries, err := LoadItineraries(settings.FilePath)

	if err != nil {
		t.Fatal("Error %s", err)
	}

	t.Run("It has connection between 114 & 142", func(tt *testing.T) {
		if !itineraries.Connected("114", "142") {
			tt.Errorf("Union between 114 & 142 not found")
		}
	})

	t.Run("It has no connection between 3 & 6", func(tt *testing.T) {
		if itineraries.Connected("3", "6") {
			tt.Errorf("Union between 3 and 6 not found")
		}
	})
}
