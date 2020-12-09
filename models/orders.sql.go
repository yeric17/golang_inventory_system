package models

import "github.com/yeric17/inventory-system/common/responses"

type Order struct {
	responses.EncodeJSON
	ID          uint64  `json:"id"`
	ClientID    uint64  `json:"client_id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Quantity    uint32  `json:"quantity"`
	Total       float64 `json:"total"`
}

type OrderCreate struct {
	responses.DecodeJSON
	ClientID    uint64 `json:"client_id"`
	ProductID   uint64 `json:"product_id"`
	Description uint32 `json:"description"`
	Quantity    uint32 `json:"quantity"`
}

type OrderData struct {
	responses.EncodeJSON
	Orders []Order `json:"data"`
}
