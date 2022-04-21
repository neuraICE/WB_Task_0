package main

type Model struct {
	OrderUid          string   `json:"order_uid" :"order_uid"`
	TrackNumber       string   `json:"track_number" :"track_number"`
	Entry             string   `json:"entry" :"entry"`
	Delivery          delivery `json:"delivery" :"delivery"`
	Payment           payment  `json:"payment" :"payment"`
	Items             []items  `json:"items" :"items"`
	Locale            string   `json:"locale" :"locale"`
	InternalSignature string   `json:"internal_signature" :"internal_signature"`
	CustomerId        string   `json:"customer_id" :"customer_id"`
	DeliveryService   string   `json:"delivery_service" :"delivery_service"`
	Shardkey          string   `json:"shardkey" :"shardkey"`
	SmId              int64    `json:"sm_id" :"sm_id"`
	DateCreated       string   `json:"date_created" :"date_created"`
	OofShard          string   `json:"oof_shard" :"oof_shard"`
}

type delivery struct {
	Name    string `json:"name" :"name"`
	Phone   string `json:"phone" :"phone"`
	Zip     string `json:"zip" :"zip"`
	City    string `json:"city" :"city"`
	Address string `json:"address" :"address"`
	Region  string `json:"region" :"region"`
	Email   string `json:"email" :"email"`
}

type payment struct {
	Transaction  string `json:"transaction" :"transaction"`
	RequestId    string `json:"request_id" :"request_id"`
	Currency     string `json:"currency" :"currency"`
	Provider     string `json:"provider" :"provider"`
	Amount       int64  `json:"amount" :"amount"`
	PaymentDt    int64  `json:"payment_dt" :"payment_dt"`
	Bank         string `json:"bank" :"bank"`
	DeliveryCost int64  `json:"delivery_cost" :"delivery_cost"`
	GoodsTotal   int64  `json:"goods_total" :"goods_total"`
	CustomFee    int64  `json:"custom_fee" :"custom_fee"`
}

type items struct {
	ChrtId      int64  `json:"chrt_id" :"chrt_id"`
	TrackNumber string `json:"track_number" :"track_number"`
	Price       int64  `json:"price" :"price"`
	rid         string `json:"rid" :"rid"`
	name        string `json:"name" :"name"`
	sale        int64  `json:"sale" :"sale"`
	size        string `json:"size" :"size"`
	totalPrice  int64  `json:"total_price" :"total_price"`
	nmId        int64  `json:"nm_id" :"nm_id"`
	brand       string `json:"brand" :"brand"`
	Status      int64  `json:"status" :"status"`
}

