package db

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/go-sqlite"
)

var (
	_db *sql.DB
	err error
)

const (
	DSN = "./db/order.db?_pragma=foreign_keys(1)"
)

func NewDBOrGet() (*sql.DB, error) {
	if nil != _db {
		return _db, nil
	}

	log.Println("[DB] trying to open database connection")
	_db, err = sql.Open("sqlite", DSN)
	if nil != err {
		log.Fatal(err)

		return nil, err
	}

	// schema init
	schemaInit(_db)

	return _db, nil
}

func schemaInit(db *sql.DB) {
	orderSchemaInit(db)
	itemSchemaInit(db)
}

func orderSchemaInit(db *sql.DB) {

	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS "order" (
            ID INTEGER PRIMARY KEY AUTOINCREMENT,
            CUSTOMER_NAME VARCHAR(50) NOT NULL,
            ORDERED_AT DATETIME,
            CREATED_AT DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
            UPDATED_AT DATETIME
        )`)
	if nil != err {
		log.Fatal(err)
	}
}

func itemSchemaInit(db *sql.DB) {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS item (
            ID INTEGER PRIMARY KEY AUTOINCREMENT,
            NAME VARCHAR(50) NOT NULL,
            DESCRIPTION VARCHAR(255),
            QUANTITY INTEGER NOT NULL DEFAULT 0,
            ORDER_ID INTEGER REFERENCES "order" (ID) ON DELETE CASCADE,
            CREATED_AT DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
            UPDATED_AT DATETIME
        )`)
	if nil != err {
		log.Fatal(err)
	}

}
