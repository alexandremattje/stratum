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

	saveBid(bidRequest)

	GetStats(w, r)

	// respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

func saveBid(bidRequest BidRequest) {
	allBids.Total_bids++

	var bid = Bid{Client_id: bidRequest.Client_id, Price: bidrequest.Price, Time_stamp: 9}

	var idx = -1
	for i := range allBids.Items {
		if allBids.Items[i].Item_id == bidRequest.Item_id {
			idx = i
			break
		}
	}

	var item Item

	if idx == -1 {
		item = Item{Item_id: bidRequest.Item_id, Hits: 1, Best: Bid}
		allBids.Items = append(allBids.Items, item)
	}

}

func GetStats(w http.ResponseWriter, r *http.Request) {
	RespondwithJSON(w, http.StatusOK, allBids)
}
