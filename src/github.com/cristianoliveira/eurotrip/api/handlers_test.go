package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cristianoliveira/eurotrip/common"
)

func TestHandlerIndex(t *testing.T) {
	repo, err := LoadItineraries(common.Settings().FilePath)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleIndex(repo))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHandleBusRouteSearch(t *testing.T) {
	repo, err := LoadItineraries(common.Settings().FilePath)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/api/direct?dep_sid=114&arr_sid=152", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleBusRouteSearch(repo))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
