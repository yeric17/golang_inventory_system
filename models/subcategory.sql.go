package models

import "github.com/yeric17/inventory-system/common/responses"

type SubCategory struct {
	responses.EncodeJSON
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type SubCategoryData struct {
	responses.EncodeJSON
	Data []SubCategory `json:"data"`
}
