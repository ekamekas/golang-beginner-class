package main

import (
    "net/http"
    "webserver-http/auth"
    "webserver-http/user"
)

const (
    SERVER_ADDRESS = ":8080"
)

func main() {
    auth.RouteInit();
    user.RouteInit();

    fs := http.FileServer(http.Dir("public"))
    http.Handle("/public/", http.StripPrefix("/public/", fs))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            http.NotFound(w, r)

            return
        }

        http.ServeFile(w, r, "public/index.html") 
    })
    
    http.ListenAndServe(SERVER_ADDRESS, nil)
}
