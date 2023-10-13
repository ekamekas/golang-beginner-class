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
    DSN = "./db/product.db?_pragma=foreign_keys(1)"
)

func GetOrInit() (*sql.DB, error) {
    if nil != _db {
        return _db, nil
    }

    log.Println("[DB] trying to open database connection")
    _db, err = sql.Open("sqlite", DSN)
    if nil != err {
        log.Fatal(err)

        return nil, err
    }
    // initialize database schemas
    ProductSchemaInit(_db)
    VariantSchemaInit(_db)

    // test
    
    return _db, nil
}
