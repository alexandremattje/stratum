package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	. "../model"
	. "../utils"
)

var allBids Stats

func LoadData() {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &allBids)

	defer jsonFile.Close()
}

func GetStats(w http.ResponseWriter, r *http.Request) {
	RespondwithJSON(w, http.StatusOK, allBids)
}

func CreateBid(w http.ResponseWriter, r *http.Request) {
	bidRequest := BidRequest{}
	fmt.Println(r.Body)
	json.NewDecoder(r.Body).Decode(&bidRequest)

	saveBid(bidRequest)

	writeData()

	GetStats(w, r)
}

func writeData() {
	file, _ := json.MarshalIndent(allBids, "", " ")
	_ = ioutil.WriteFile("data.json", file, 0644)
}

func saveBid(bidRequest BidRequest) {
	allBids.Total_bids++

	var bid = Bid{Client_id: bidRequest.Client_id, Price: bidRequest.Price, Time_stamp: time.Now().Unix()}

	var idx = -1
	for i := range allBids.Items {
		if allBids.Items[i].Item_id == bidRequest.Item_id {
			idx = i
			break
		}
	}

	var item Item

	fmt.Println(bidRequest)
	if idx == -1 {
		item = Item{Item_id: bidRequest.Item_id, Hits: 1, Best: bid}
		fmt.Println(allBids)
		allBids.Items = append(allBids.Items, item)
		fmt.Println(allBids)
	} else {
		if isBetterThanBest(allBids.Items[idx].Best, bid) {
			allBids.Items[idx].Best = bid
			allBids.Items[idx].Hits++
		}
	}

}

func isBetterThanBest(best Bid, bid Bid) bool {
	return bid.Price > best.Price
}
