package models

import "github.com/yeric17/inventory-system/common/responses"

type Brand struct {
	responses.EncodeJSON
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type BrandData struct {
	responses.EncodeJSON
	Data []Brand `json:"data"`
}
