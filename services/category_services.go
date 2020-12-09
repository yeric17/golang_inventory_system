package services

import (
	"fmt"

	"github.com/yeric17/inventory-system/db"
	"github.com/yeric17/inventory-system/models"
)

var (
	CategoryServices = categoryServices{}
)

type categoryServices struct{}

func (categoryServices) GetByName(categoryName string) (category models.Category, err1 error) {
	var query string
	query = fmt.Sprintf(`
		SELECT id, name
		FROM categories
		WHERE name = $1`,
	)

	stmt, err := db.DBClient.Prepare(query)

	if err != nil {
		return category, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(categoryName)

	if err := row.Scan(
		&category.ID,
		&category.Name,
	); err != nil {
		return category, err
	}

	return category, nil
}
