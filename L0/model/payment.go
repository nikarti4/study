package model

type payment_t struct {
	Transaction   string  `json:"transaction"`
	Request_id    string  `json:"request_id"`
	Currency      string  `json:"currency"`
	Provider      string  `json:"provider"`
	Amount        uint    `json:"amount"`
	Payment_dt    uint    `json:"payment_dt"`
	Bank          string  `json:"bank"`
	Delivery_cost float64 `json:"delivery_cost"`
	Goods_total   uint    `json:"goods_total"`
	Custom_fee    uint    `json:"custom_fee"`
}
