package services

import (
	"fmt"

	"github.com/yeric17/inventory-system/db"
	"github.com/yeric17/inventory-system/models"
)

var (
	SubCategoryServices = subcategoryServices{}
)

type subcategoryServices struct{}

func (subcategoryServices) GetByName(subcategoryName string) (subcategory models.SubCategory, err1 error) {
	var query string
	query = fmt.Sprintf(`
		SELECT id, name
		FROM subcategories
		WHERE name = $1`,
	)

	stmt, err := db.DBClient.Prepare(query)

	if err != nil {
		return subcategory, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(subcategoryName)

	if err := row.Scan(
		&subcategory.ID,
		&subcategory.Name,
	); err != nil {
		return subcategory, err
	}

	return subcategory, nil
}
