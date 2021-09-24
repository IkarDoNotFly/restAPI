package store

import (
	"encoding/json"
	"github.com/gorilla/websocket"
)

type PairRepo struct {
	store *Store
}

type Request struct{
	Method string `json:"method"`
	Params [1]string `json:"symbol"`
	ID int `json:"id"`
}
//endpoint wss://stream.binance.com:9443
func something(symbol [1]string){
	conn,res,err:=websocket.DefaultDialer.Dial("wss://stream.binancefuture.com",nil)
	test:=Request{
		Method:"SUBSCRIBE",
		Params:symbol,
		ID: 1,
	}
}