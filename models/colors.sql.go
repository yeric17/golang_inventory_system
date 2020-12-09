package models

import "github.com/yeric17/inventory-system/common/responses"

type Color struct {
	responses.EncodeJSON
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Hex  string `json:"hex"`
}

type ColorData struct {
	responses.EncodeJSON
	Data []Color `json:"data"`
}
