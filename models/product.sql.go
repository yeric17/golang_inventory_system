package models

import (
	"database/sql"

	"github.com/yeric17/inventory-system/common/convert"
	"github.com/yeric17/inventory-system/common/responses"
)

//Product es la estructura de un producto a imagen de la base de datos
type Product struct {
	responses.EncodeJSON
	ID                 uint64  `json:"id"`
	Name               string  `json:"name"`
	Price              float64 `json:"price"`
	ReorderLevel       uint32  `json:"reorder_level"`
	Size               string  `json:"size"`
	Description        string  `json:"description"`
	ImageURL           string  `json:"image_url"`
	HasDiscount        bool    `json:"has_discount"`
	DiscountPercentage float64 `json:"discount_percentage"`
	Brand              string  `json:"brand"`
	Category           string  `json:"category"`
	SubCategory        string  `json:"subcategory,omitempty"`
	Color              string  `json:"color"`
	Hex                string  `json:"hex"`
	Quantity           uint32  `json:"quantity"`
	Status             string  `json:"status"`
}

//ProductSQL sirve para controla los posibles null de Product
type ProductSQL struct {
	ID                 uint64
	Name               string
	Price              float64
	ReorderLevel       uint32
	Size               string
	Description        sql.NullString
	ImageURL           sql.NullString
	HasDiscount        int8
	DiscountPercentage float64
	Brand              string
	Category           string
	SubCategory        sql.NullString
	Color              string
	Hex                string
	Quantity           uint32
	Status             string
}

//ProductUpdate sirve para controla los posibles null de Product
type ProductUpdate struct {
	responses.DecodeJSON
	ID                 uint64  `json:"id"`
	Name               string  `json:"name"`
	Price              float64 `json:"price"`
	ReorderLevel       uint32  `json:"reorder_level"`
	Size               string  `json:"size"`
	Description        string  `json:"description"`
	ImageURL           string  `json:"image_url"`
	HasDiscount        bool    `json:"has_discount"`
	DiscountPercentage float64 `json:"discount_percentage"`
	Brand              string  `json:"brand"`
	Category           string  `json:"category"`
	SubCategory        string  `json:"subcategory"`
	Color              string  `json:"color"`
	Status             string  `json:"status"`
}

//ProductData contiene un slice de Product
type ProductData struct {
	responses.EncodeJSON
	Products []Product `json:"products"`
}

type ProductFilters struct {
	Name               string
	MaxPrice           float64
	MinPrice           float64
	Price              float64
	MinQuantity        int
	MaxQuantity        int
	Description        string
	HasDiscount        bool
	DiscountPercentage float64
	Brand              string
	Category           string
	SubCategory        string
	Color              string
	Status             string
}

//ToProduct transforma un ProductNull ha Product
func (pn *ProductSQL) ToProduct() (prod Product) {
	prod.ID = pn.ID
	prod.Name = pn.Name
	prod.Price = pn.Price
	prod.ReorderLevel = pn.ReorderLevel
	prod.Size = pn.Size
	prod.Description = pn.Description.String
	prod.ImageURL = pn.ImageURL.String
	prod.HasDiscount = convert.IntToBool(pn.HasDiscount)
	prod.DiscountPercentage = pn.DiscountPercentage
	prod.Brand = pn.Brand
	prod.Category = pn.Category
	prod.SubCategory = pn.SubCategory.String
	prod.Color = pn.Color
	prod.Hex = pn.Hex
	prod.Quantity = pn.Quantity
	prod.Status = pn.Status
	return
}

//FromProduct transforma un Product ha ProductSQL
func (pn *Product) FromProduct() (prod ProductSQL) {
	prod.ID = pn.ID
	prod.Name = pn.Name
	prod.Price = pn.Price
	prod.ReorderLevel = pn.ReorderLevel
	prod.Size = pn.Size
	if pn.Description == "" {
		prod.Description.Valid = false
	} else {
		prod.Description.String = pn.Description
	}
	if pn.ImageURL == "" {
		prod.ImageURL.Valid = false
	} else {
		prod.ImageURL.String = pn.ImageURL
	}
	prod.HasDiscount = convert.BoolToInt(pn.HasDiscount)
	prod.DiscountPercentage = pn.DiscountPercentage
	// prod.CreateAt = pn.CreateAt
	// prod.UpdateAt = pn.UpdateAt
	if pn.SubCategory == "" {
		prod.SubCategory.Valid = false
	} else {
		prod.SubCategory.String = pn.SubCategory
	}
	prod.Color = pn.Color
	prod.Hex = pn.Hex
	prod.Quantity = pn.Quantity
	prod.Status = pn.Status
	return
}
