package main

import (
    "log"

    // data source adapter
    "webserver-http/db"
)

func main() {
    // database initialization
    mDB, err := db.NewDBOrGet()
    if nil != err {
        log.Fatal(err)
    }
}
