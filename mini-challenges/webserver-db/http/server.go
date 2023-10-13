package http

import (
    "net/http"
    "log"
)

const (
    SERVER_ADDRESS = ":8080"
)

func ServerRun() {
    // controller initialize
    ProductControllerInit()
    VariantControllerInit()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println(r)
        http.NotFound(w, r)

        return
    })
    
    log.Println("Running http server at", SERVER_ADDRESS)
    http.ListenAndServe(SERVER_ADDRESS, nil)
}
