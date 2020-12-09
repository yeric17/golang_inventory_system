package services

import (
	"fmt"

	"github.com/yeric17/inventory-system/db"
	"github.com/yeric17/inventory-system/models"
)

var (
	ColorServices = colorServices{}
)

type colorServices struct{}

func (colorServices) GetByName(colorName string) (color models.Color, err1 error) {
	var query string
	query = fmt.Sprintf(`
		SELECT id, name, hex
		FROM colors
		WHERE name = $1`,
	)

	stmt, err := db.DBClient.Prepare(query)

	if err != nil {
		return color, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(colorName)

	if err := row.Scan(
		&color.ID,
		&color.Name,
		&color.Hex,
	); err != nil {
		return color, err
	}

	return color, nil
}
