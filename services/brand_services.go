package services

import (
	"fmt"

	"github.com/yeric17/inventory-system/db"
	"github.com/yeric17/inventory-system/models"
)

var (
	BrandServices = brandServices{}
)

type brandServices struct{}

func (brandServices) GetByName(brandName string) (brand models.Brand, err1 error) {
	var query string
	query = fmt.Sprintf(`
		SELECT id, name
		FROM brands
		WHERE name = $1`,
	)

	stmt, err := db.DBClient.Prepare(query)

	if err != nil {
		return brand, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(brandName)

	if err := row.Scan(
		&brand.ID,
		&brand.Name,
	); err != nil {
		return brand, err
	}

	return brand, nil
}
