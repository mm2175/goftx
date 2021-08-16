package models

import (
	"fmt"
	"time"
)

type CreateWithdrawPayload struct {
	Coin    string  `json:"coin"`
	Size    float64 `json:"size"`
	Address string  `json:"address"`
	Tag     string  `json:"tag,omitempty"`
	Method  string  `json:"method,omitempty"`
}

type Withdraw struct {
	ID      int64     `json:"id"`
	Coin    string    `json:"coin"`
	Address string    `json:"address"`
	Tag     string    `json:"tag"`
	Fee     float64   `json:"fee"`
	Size    float64   `json:"size"`
	Status  string    `json:"status"`
	TxID    string    `json:"txid"`
	Time    time.Time `json:"time"`
}

func (w Withdraw) String() string {
	return fmt.Sprintf("ftx.Withdraw{id: %v, coin: %v, address: %v, tag: %v, fee: %v, size: %v, status: %v, time: %v, txid: %v}",
		w.ID, w.Coin, w.Address, w.Tag, w.Fee, w.Size, w.Status, w.Time, w.TxID)
}
