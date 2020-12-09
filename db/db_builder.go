package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/yeric17/inventory-system/config"
)

type dbConnection interface {
	getDNS() string
	connect(*sql.DB)
}

type postgresDB struct{}

func (postgresDB) getDNS() string {
	var connString string
	connString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.DBname)
	return connString
}
func (p postgresDB) connect(sqlDB *sql.DB) {
	var err error
	sqlDB, err = sql.Open(config.Driver, p.getDNS())
	if err != nil {
		panic(err)
	}

	err = sqlDB.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("Conectado a la base de datos postgres")
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)
}

func getInstance(driver string) dbConnection {
	switch driver {
	case "postgres":
		return postgresDB{}
	default:
		return nil
	}
}

func startConnection(driver string, sq *sql.DB) {
	var conn dbConnection
	conn = getInstance(driver)
	//TODO solucionar identificación de driver no soportado
	if conn == nil {
		fmt.Println("Driver no soportado aun")
		return
	}
	conn.connect(sq)
}
