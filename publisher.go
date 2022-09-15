package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

var tests1 = `{
  "order_uid": "b111feb7b2b84b6test",
  "track_number": "TRACK1",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "transaction": "b111feb7b2b84b6test",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "TRACK1",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}`

var wrong = `{
  "order_uid": "b563feb7b2b84b6test",
  "track_number": "TRACK2",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  }
}`

var tests = `{
  "order_uid": "b555feb7b2b84b6test",
  "track_number": "TRACK3",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "transaction": "b555feb7b2b84b6test",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
"items": [
    {
    "chrt_id": 1,
    "track_number": "TRACK3",
    "price": 1000,
    "rid": "ab4219087a764ae0btest",
    "name": "bike",
    "sale": 10,
    "size": "0",
    "total_price": 900,
    "nm_id": 2383124,
    "brand": "Yamaha",
    "status": 100
    },
    {
    "chrt_id": 9999,
    "track_number": "TRACK3",
    "price": 500,
    "rid": "ab4219087a764ae0btest",
    "name": "Helmet",
    "sale": 50,
    "size": "0",
    "total_price": 250,
    "nm_id": 2389212,
    "brand": "Kymco",
    "status": 201
    },
    {
    "chrt_id": 32,
    "track_number": "TRACK3",
    "price": 123,
    "rid": "ab4219087a764ae0btest",
    "name": "T-shirt",
    "sale": 30,
    "size": "0",
    "total_price": 317,
    "nm_id": 2389212,
    "brand": "Nike",
    "status": 202
    },
    {
    "chrt_id": 23,
    "track_number": "TRACK3",
    "price": 999,
    "rid": "ab4219087a764ae0btest",
    "name": "Pants",
    "sale": 30,
    "size": "0",
    "total_price": 666,
    "nm_id": 2389212,
    "brand": "Vivienne Sabo",
    "status": 202
    }
],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}`

func main() {
	conn, err := stan.Connect("test-cluster", "test-publisher")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	err = conn.Publish("service", []byte(tests))
	fmt.Println("Right message")
	err = conn.Publish("service", []byte(tests1))
	fmt.Println("Right message")
	err = conn.Publish("service", []byte("Hello"))
	fmt.Println("Wrong message")
	err = conn.Publish("service", []byte(wrong))
	fmt.Println("Wrong message")
}
