package db

import (
	"database/sql"

	//github.com/lib/pq driver de postgres
	_ "github.com/lib/pq"
	"github.com/yeric17/inventory-system/config"
)

var (
	//DBClient es un cliente de SQL
	DBClient *sql.DB
)

func init() {
	startConnection(config.Driver, DBClient)
}
