package controller

import (
	"encoding/json"
	"net/http"

	. "../model"
	. "../utils"
)

var allBids Stats

func CreateBid(w http.ResponseWriter, r *http.Request) {
	bidRequest := BidRequest{}
	json.NewDecoder(r.Body).Decode(&bidRequest)

	allBids.Total_bids++

	GetStats(w, r)

	// respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

func GetStats(w http.ResponseWriter, r *http.Request) {
	RespondwithJSON(w, http.StatusOK, allBids)
}
