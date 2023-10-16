package http

import (
	"encoding/json"
	"log"
	"net/http"
	"webserver-http/domain"
)

func NewOrderRoute(repository domain.OrderRepository) {
	controller := domain.NewOrderController(repository)

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			var requestBody domain.Order
			err := json.NewDecoder(r.Body).Decode(&requestBody)
			if nil != err {
				log.Println("[ORDER] failed to process request", err)
				handleGeneric(w, r)

				break
			}

			responseBody := controller.Create(&requestBody)

			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(responseBody)
			if nil != err {
				log.Println("[ORDER] failed to process request, err")
				handleGeneric(w, r)
			}

			break
		case "GET":
			responseBody := controller.Get()

			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(responseBody)
			if nil != err {
				log.Println("[ORDER] failed to process request, err")
				handleGeneric(w, r)
			}

			break
		default:
			handleMethodNotAllowed(w, r)

			break
		}
	})
}
