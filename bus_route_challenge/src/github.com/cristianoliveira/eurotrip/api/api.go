package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cristianoliveira/eurotrip/common"
)

func Serve(settings *common.Setting) {
	address := "0.0.0.0:" + settings.Port

	repo, err := LoadItineraries(settings.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data source: " + settings.FilePath)
	fmt.Println("")
	fmt.Println("Running on: " + address)
	fmt.Println("Endpoints:")
	fmt.Println("/api/")
	fmt.Println("/api/direct?dep_sid={}&arr_sid={}")

	http.HandleFunc("/api/", HandleIndex(repo))
	http.HandleFunc("/api/direct", HandleBusRouteSearch(repo))

	log.Fatal(http.ListenAndServe(address, nil))
}
