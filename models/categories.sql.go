package models

import "github.com/yeric17/inventory-system/common/responses"

type Category struct {
	responses.EncodeJSON
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
type CategoryData struct {
	responses.EncodeJSON
	Data []Category `json:"data"`
}
