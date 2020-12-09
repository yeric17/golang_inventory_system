package services

import (
	"fmt"

	"github.com/yeric17/inventory-system/db"
	"github.com/yeric17/inventory-system/models"
)

var (
	StatusServices = statusServices{}
)

type statusServices struct{}

func (statusServices) GetByName(statusName string) (status models.Status, err1 error) {
	var query string
	query = fmt.Sprintf(`
		SELECT id, name
		FROM status
		WHERE name = $1`,
	)

	stmt, err := db.DBClient.Prepare(query)

	if err != nil {
		return status, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(statusName)

	if err := row.Scan(
		&status.ID,
		&status.Name,
	); err != nil {
		return status, err
	}

	return status, nil
}
