package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/yeric17/inventory-system/common/convert"
	"github.com/yeric17/inventory-system/common/tools"
	"github.com/yeric17/inventory-system/db"
	"github.com/yeric17/inventory-system/models"
)

var (
	ProductServices = productServices{}
)

const fullProductQuery = `
SELECT 
products.id, 
products.name,
products.price,
products.reorder_level,
sizes.name size,
products.description,
images.url image_url,
products.has_discount,
products.discount_percentage,
brands.name brand,
categories.name category,
subcategories.name subcategory,
colors.name color,
colors.hex,
stock.quantity,
status.name status
FROM products
LEFT JOIN sizes ON products.size_id = sizes.id
LEFT JOIN brands ON products.brand_id = brands.id
LEFT JOIN categories ON products.category_id = categories.id
LEFT JOIN subcategories ON products.subcategory_id = subcategories.id
LEFT JOIN colors ON products.color_id = colors.id
LEFT JOIN stock ON products.id = stock.product_id
LEFT JOIN status ON products.status_id = status.id
LEFT JOIN images ON products.image_id = images.id
`

type productServices struct{}

func (productServices) GetAll(limitRows int32, page int32) (data models.ProductData, err1 error) {
	var query string
	query = fmt.Sprintf(`%s LIMIT %d OFFSET %d`, fullProductQuery, limitRows, page-1)

	stmt, err := db.DBClient.Prepare(query)

	if err != nil {
		return data, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		var prod models.ProductSQL

		if err := rows.Scan(
			&prod.ID,
			&prod.Name,
			&prod.Price,
			&prod.ReorderLevel,
			&prod.Size,
			&prod.Description,
			&prod.ImageURL,
			&prod.HasDiscount,
			&prod.DiscountPercentage,
			&prod.Brand,
			&prod.Category,
			&prod.SubCategory,
			&prod.Color,
			&prod.Hex,
			&prod.Quantity,
			&prod.Status,
		); err != nil {
			return data, err
		}
		data.Products = append(data.Products, prod.ToProduct())
	}

	return data, nil
}

func (productServices) GetByName(productName string) (prod models.Product, err1 error) {
	var query string
	query = fmt.Sprintf(`%s WHERE products.name = %s`, fullProductQuery, productName)

	stmt, err := db.DBClient.Prepare(query)

	if err != nil {
		return prod, err
	}

	defer stmt.Close()

	row := stmt.QueryRow()
	var prodSQL models.ProductSQL

	if err := row.Scan(
		&prodSQL.ID,
		&prodSQL.Name,
		&prodSQL.Price,
		&prodSQL.ReorderLevel,
		&prodSQL.Size,
		&prodSQL.Description,
		&prodSQL.ImageURL,
		&prodSQL.HasDiscount,
		&prodSQL.DiscountPercentage,
		&prodSQL.Brand,
		&prodSQL.Category,
		&prodSQL.SubCategory,
		&prodSQL.Color,
		&prodSQL.Hex,
		&prodSQL.Quantity,
		&prodSQL.Status,
	); err != nil {
		return prod, err
	}
	prod = prodSQL.ToProduct()
	return prod, nil
}

func (productServices) Get(productID uint64) (prod models.Product, err1 error) {
	var query string
	query = fmt.Sprintf(`%s WHERE products.id = %d`, fullProductQuery, productID)

	stmt, err := db.DBClient.Prepare(query)

	if err != nil {
		return prod, err
	}

	defer stmt.Close()

	row := stmt.QueryRow()
	var prodSQL models.ProductSQL

	if err := row.Scan(
		&prodSQL.ID,
		&prodSQL.Name,
		&prodSQL.Price,
		&prodSQL.ReorderLevel,
		&prodSQL.Size,
		&prodSQL.Description,
		&prodSQL.ImageURL,
		&prodSQL.HasDiscount,
		&prodSQL.DiscountPercentage,
		&prodSQL.Brand,
		&prodSQL.Category,
		&prodSQL.SubCategory,
		&prodSQL.Color,
		&prodSQL.Hex,
		&prodSQL.Quantity,
		&prodSQL.Status,
	); err != nil {
		return prod, err
	}
	prod = prodSQL.ToProduct()
	return prod, nil
}

func (productServices) GetFilter(limitRows int32, page int32, filter models.ProductFilters) (models.ProductData, error) {
	var mapQuery []string
	var numOfRows int32
	var products models.ProductData

	if filter.Name != "" {
		mapQuery = append(mapQuery, fmt.Sprintf("(%s)", tools.StringToQueryLocate(filter.Name, "name")))
	}
	if filter.Description != "" {
		mapQuery = append(mapQuery, fmt.Sprintf("(%s)", tools.StringToQueryLocate(filter.Description, "description")))
	}

	selectString := strings.Join(mapQuery[:], " AND ")

	if selectString == "" {
		return models.ProductData{}, errors.New("[Error] The filter can't be empty")
	}

	query := fmt.Sprintf("%s WHERE %s LIMIT %d OFFSET %d", fullProductQuery, selectString, limitRows, page-1)
	query2 := fmt.Sprintf("SELECT FOUND_ROWS() AS num_rows FROM products WHERE %s", selectString)

	stmtRows, err := db.DBClient.Prepare(query2)
	if err != nil {
		return models.ProductData{}, err
	}
	defer stmtRows.Close()

	rowsCount, err := stmtRows.Query()

	if err != nil {
		return models.ProductData{}, err
	}
	defer rowsCount.Close()

	for rowsCount.Next() {
		numOfRows++
	}

	stmt, err := db.DBClient.Prepare(query)
	if err != nil {
		return models.ProductData{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return models.ProductData{}, err
	}
	defer rows.Close()

	for rows.Next() {

		prod := models.ProductSQL{}

		if err := rows.Scan(
			&prod.ID,
			&prod.Name,
			&prod.Price,
			&prod.ReorderLevel,
			&prod.Size,
			&prod.Description,
			&prod.ImageURL,
			&prod.HasDiscount,
			&prod.DiscountPercentage,
			&prod.Brand,
			&prod.Category,
			&prod.SubCategory,
			&prod.Color,
			&prod.Hex,
			&prod.Quantity,
			&prod.Status,
		); err != nil {
			return models.ProductData{}, err
		}

		products.Products = append(products.Products, prod.ToProduct())
	}

	if len(products.Products) == 0 {
		return models.ProductData{}, errors.New("[Error] No se logro encontrar ningún producto")
	}

	return products, nil
}

//Update pend comment
func (productServices) Update(prodUpdate models.ProductUpdate) (models.Product, error) {
	if prodUpdate.ID == 0 {
		return models.Product{}, fmt.Errorf("product.ProductID can't be null for update: %+v", prodUpdate)
	}
	var numFields int
	var values []interface{}
	var fields []string

	prod, err := ProductServices.Get(prodUpdate.ID)

	numFields = 0

	if prodUpdate.Name != "" && prodUpdate.Name != prod.Name {
		numFields++
		values = append(values, prodUpdate.Name)
		fields = append(fields, fmt.Sprintf("SET name = $%d", numFields))
	}

	if prodUpdate.Price != 0 && prodUpdate.Price != prod.Price {
		numFields++
		values = append(values, prodUpdate.Price)
		fields = append(fields, "SET price = $%d")
	}

	if prodUpdate.ReorderLevel != 0 && prodUpdate.ReorderLevel != prod.ReorderLevel {
		numFields++
		values = append(values, prodUpdate.ReorderLevel)
		fields = append(fields, "SET = reorder_level = $%d")
	}

	if prodUpdate.Size != "" && prodUpdate.Size != prod.Size {
		numFields++
		values = append(values, prodUpdate.Size)
		fields = append(fields, "SET size = $%d")
	}

	if prodUpdate.Description != "" && prodUpdate.Description != prod.Description {
		numFields++
		values = append(values, prodUpdate.Description)
		fields = append(fields, "SET description = $%d")
	}

	if prodUpdate.ImageURL != "" && prodUpdate.ImageURL != prod.ImageURL {
		numFields++
		values = append(values, prodUpdate.ImageURL)
		fields = append(fields, "SET image_url = $%d")
	}

	if prodUpdate.HasDiscount != prod.HasDiscount && prodUpdate.DiscountPercentage > 0 {
		numFields++
		values = append(values, convert.BoolToInt(prodUpdate.HasDiscount))
		fields = append(fields, "SET has_discount = $%d")
	}

	if prodUpdate.DiscountPercentage != 0 && prodUpdate.HasDiscount && prodUpdate.DiscountPercentage != prod.DiscountPercentage {
		numFields++
		values = append(values, prodUpdate.DiscountPercentage)
		fields = append(fields, "SET discount_percentage = $%d")
	}
	if prodUpdate.Brand != "" && prodUpdate.Brand != prod.Brand {
		numFields++
		brand, err := BrandServices.GetByName(prodUpdate.Brand)
		if err != nil {
			return models.Product{}, err
		}
		values = append(values, brand.ID)
		fields = append(fields, "SET brand_id = $%d")
	}
	if prodUpdate.Category != "" && prodUpdate.Category != prod.Category {
		numFields++
		category, err := CategoryServices.GetByName(prodUpdate.Category)
		if err != nil {
			return models.Product{}, err
		}
		values = append(values, category.ID)
		fields = append(fields, "SET category_id = $%d")
	}

	if prodUpdate.SubCategory != "" && prodUpdate.SubCategory != prod.SubCategory {
		numFields++
		subcategory, err := SubCategoryServices.GetByName(prodUpdate.SubCategory)
		if err != nil {
			return models.Product{}, err
		}
		values = append(values, subcategory.ID)
		fields = append(fields, "SET subcategory_id = $%d")
	}

	if prodUpdate.Color != "" && prodUpdate.Color != prod.Color {
		numFields++
		color, err := ColorServices.GetByName(prodUpdate.Color)
		if err != nil {
			return models.Product{}, err
		}
		values = append(values, color.ID)
		fields = append(fields, "SET color_id = $%d")
	}

	if prodUpdate.Status != "" && prodUpdate.Status != prod.Status {
		numFields++
		status, err := StatusServices.GetByName(prodUpdate.Status)
		if err != nil {
			return models.Product{}, err
		}
		values = append(values, status.ID)
		fields = append(fields, "SET status_id = $%d")
	}

	if len(values) < 1 {
		return prod, errors.New("[Error] No no existen valores en el body para actualizar")
	}

	var query string

	query = fmt.Sprintf("UPDATE products %s  WHERE products.id = %d;", strings.Join(fields[:], ", "), prodUpdate.ID)

	fmt.Println(query)

	stmt, err := db.DBClient.Prepare(query)

	if err != nil {
		return models.Product{}, err
	}
	defer stmt.Close()

	results, err := stmt.Exec(values...)

	if err != nil {
		return models.Product{}, err
	}

	rowsAffected, err := results.RowsAffected()

	if err != nil {
		return models.Product{}, err
	}

	if rowsAffected != 1 {
		return models.Product{}, fmt.Errorf("[Error] Se esperaba 1 fila afectada, encambio se afectaron: [%d]", rowsAffected)
	}

	prod, err = ProductServices.Get(prod.ID)
	if err != nil {
		return models.Product{}, err
	}

	return prod, nil
}

//Delete pend comment
func (productServices) Delete(productID uint64) (models.Product, error) {
	stmt, err := db.DBClient.Prepare("DELETE FROM products WHERE products.id = $1")
	if err != nil {
		return models.Product{}, err
	}
	defer stmt.Close()

	prod, _ := ProductServices.Get(productID)

	result, err := stmt.Exec(productID)
	if err != nil {
		return models.Product{}, err
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return models.Product{}, err
	}

	if rowsAffected != 1 {
		return models.Product{}, fmt.Errorf("[Error] Se esperaba 1 fila afectada, encambio se afectaron: [%d]", rowsAffected)
	}

	return prod, nil
}
