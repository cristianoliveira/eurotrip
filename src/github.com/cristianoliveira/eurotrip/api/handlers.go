package api

import (
	"encoding/json"
	"net/http"
)

type ResponseData struct {
	DepSid string "json:dep_sid"
	ArrSid string "json:arr_sid"
	Direct bool   "json:direct_bus_route"
}

type Handler func(w http.ResponseWriter, r *http.Request)

func HandleIndex(repo *Itinerary) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		js, err := json.Marshal(repo.Routes())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(js)
	}
}

func HandleBusRouteSearch(repo *Itinerary) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		depSid := r.URL.Query().Get("dep_sid")
		arrSid := r.URL.Query().Get("arr_sid")
		if depSid == "" || arrSid == "" {
			http.Error(w, "Missing parameters", http.StatusBadRequest)
			return
		}

		response := ResponseData{
			DepSid: depSid,
			ArrSid: arrSid,
			Direct: repo.Connected(depSid, arrSid),
		}

		js, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(js)
	}
}
