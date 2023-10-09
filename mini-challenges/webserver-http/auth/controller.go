package auth

import (
    "log"
    "net/http"
    "html/template"
    "webserver-http/user"
)

func RouteInit() {
    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
            case "GET":
                loginPage(w, r)
                break
            case "POST":
                login(w, r)
                break
            default:
                 http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)       
        }
    });
    http.HandleFunc("/not-registered", notRegisteredPage);
}

func notRegisteredPage(w http.ResponseWriter, r *http.Request) {
    if("GET" != r.Method) {
        http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
        return
    }

    http.ServeFile(w, r, "public/not-registered.html") 
}


func loginPage(w http.ResponseWriter, r *http.Request) {
    if("GET" != r.Method) {
        http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
        return
    }

    tpl, error := template.New("login.html").ParseFiles("public/login.html")
    if(nil != error) {
        log.Println(error.Error())
        http.Error(w, "Failed to process request", http.StatusInternalServerError)

        return
    }

    data := struct {
        Emails []string
    }{ user.GetEmails() }

    if error := tpl.ExecuteTemplate(w, "login.html", data); nil != error {
        log.Println(error.Error())
        http.Error(w, "Failed to process request", http.StatusInternalServerError)

        return;
    }
}

func login(w http.ResponseWriter, r *http.Request) {
    if("POST" != r.Method) {
        http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)

        return
    }
    
    // validate data
    loggedUser, error := user.GetByEmail(r.FormValue("email"))
    if nil != error {
        log.Println(error.Error())

        w.Header().Set("Location", "/not-registered")
        w.WriteHeader(http.StatusSeeOther)

        return
    }

    w.Header().Set("Location", "/profile?token=" + loggedUser.Email)
    w.WriteHeader(http.StatusSeeOther)
}
