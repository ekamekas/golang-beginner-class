package user

import (
    "log"
    "net/http"
    "html/template"
)

func RouteInit() {
    http.HandleFunc("/profile", profilePage)
}

func profilePage(w http.ResponseWriter, r *http.Request) {
    if("GET" != r.Method) {
        http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
        return
    }

    token := r.URL.Query().Get("token")
    loggedUser, error := GetByEmail(token)
    if nil != error {
        log.Println(error.Error())

        w.Header().Set("Location", "/not-registered")
        w.WriteHeader(http.StatusSeeOther)

        return
    }

    tpl, error := template.New("profile.html").ParseFiles("public/profile.html")
    if(nil != error) {
        log.Println(error.Error())
        http.Error(w, "Failed to process request", http.StatusInternalServerError)

        return
    }

    data := struct {
        User User
        Token string
    }{ loggedUser, token }


    if error := tpl.ExecuteTemplate(w, "profile.html", data); nil != error {
        log.Println(error.Error())
        http.Error(w, "Failed to process request", http.StatusInternalServerError)

        return;
    }
}
