package main

import (
	"log"

	// data source adapter
	"webserver-http/db"

	// rest adapter
	"webserver-http/http"
)

func main() {
	// database initialization
	mDB, err := db.NewDBOrGet()
	if nil != err {
		log.Fatal(err)
	}

	// repository initialization
	orderRepository := db.NewOrderRepository(mDB)

	// route initialization
	http.NewOrderRoute(orderRepository)
	http.ServerRun()
}
