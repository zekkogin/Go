package model

import (
	"time"
)

type Order struct {
	OrderUID          string    `json:"order_uid" fake:"{uuid}test"`
	TrackNumber       string    `json:"track_number" fake:"{regex:[WBILMTERACK]{14}}"`
	Entry             string    `json:"entry" fake:"{regex:[WBILMTERACK]{4}}"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Item    `json:"items" fakesize:"1,2"`
	Locale            string    `json:"locale" fake:"{randomstring:[en, ru]}"`
	InternalSignature string    `json:"internal_signature" fake:"skip"`
	CustomerID        string    `json:"customer_id" fake:"test"`
	DeliveryService   string    `json:"delivery_service" fake:"meest"`
	Shardkey          string    `json:"shardkey" fake:"{number:1,10}"`
	SmID              int64     `json:"sm_id" fake:"{number:1,100}"`
	DateCreated       time.Time `json:"date_created" fake:"{date}"`
	OofShard          string    `json:"oof_shard" fake:"1"`
}

type Delivery struct {
	Name    string `json:"name" fake:"{randomstring:[Navuhodonosor, Dima Dimovich]}"`
	Phone   string `json:"phone" fake:"{phone}"`
	Zip     string `json:"zip" fake:"{zip}"`
	City    string `json:"city" fake:"{randomstring:[Moscow, Cow, Cumbodja, {city}]}"`
	Address string `json:"address" fake:"{streetname} {streetnumber}"`
	Region  string `json:"region" fake:"{state}"`
	Email   string `json:"email" fake:"{email}"`
}

type Payment struct {
	Transaction  string `json:"transaction" fake:"{uuid}test"`
	RequestID    string `json:"request_id" fake:"skip"`
	Currency     string `json:"currency" fake:"{randomstring:[USD,RUB,SOM]}"`
	Provider     string `json:"provider" fake:"wbpay"`
	Amount       int64  `json:"amount" fake:"{number:1,100000}"`
	PaymentDt    int64  `json:"payment_dt" fake:"{number:1,1637907727}"`
	Bank         string `json:"bank" fake:"{randomstring:[alpha,sber,tinkoff,roketa]}"`
	DeliveryCost int64  `json:"delivery_cost" fake:"{number:100,1500}"`
	GoodsTotal   int64  `json:"goods_total" fake:"skip"`
	CustomFee    int64  `json:"custom_fee" fake:"skip"`
}

type Item struct {
	ChrtID      int    `json:"chrt_id" fake:"{number:1000000,9999999}"`
	TrackNumber string `json:"track_number" fake:"WBILMTESTTRACK"`
	Price       int64  `json:"price" fake:"{number:24,100000}"`
	Rid         string `json:"rid" fake:"{uuid}test"`
	Name        string `json:"name" fake:"{hipsterword}"`
	Sale        int64  `json:"sale" fake:"skip"`
	Size        string `json:"size" fake:"0"`
	TotalPrice  int64  `json:"total_price" fake:"skip"`
	NmID        int64  `json:"nm_id" fake:"{number:1000,999999}"`
	Brand       string `json:"brand" fake:"{slogan}"`
	Status      int64  `json:"status" fake:"{httpstatuscode}"`
}
