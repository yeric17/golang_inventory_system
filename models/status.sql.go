package models

import "github.com/yeric17/inventory-system/common/responses"

type Status struct {
	responses.EncodeJSON
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type StatusData struct {
	responses.EncodeJSON
	Data []Status `json:"data"`
}
