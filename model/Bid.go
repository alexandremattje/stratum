package model

type Stats struct {
	Total_bids int    `json:"total_bids,omitempty"`
	Total_hits int    `json:"total_hits,omitempty"`
	Items      []Item `json:"items,omitempty"`
}

type Item struct {
	Item_id string `json:"item_id,omitempty"`
	Hits    int    `json:"hits,omitempty"`
	Best    Bid    `json:"best_bid,omitempty"`
}

type Bid struct {
	Client_id  string  `json:"client_id,omitempty"`
	Price      float32 `json:"price,omitempty"`
	Time_stamp int64   `json:"time_stamp,omitempty"`
}

type BidRequest struct {
	Item_id   string  `json:"item_id,omitempty"`
	Price     float32 `json:"price,omitempty"`
	Client_id string  `json:"client_id,omitempty"`
}
