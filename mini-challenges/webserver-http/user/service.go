package user

import "errors"

type User struct {
    Email string
    Address string
    Job string
}

var user = map[string]User {
    "john.doe@example.com": User {
        Email: "john.doe@example.com",
        Address: "East Jakarta, Indonesia",
        Job: "Software Engineer",
    },
    "jane.smith@example.com": User {
        Email: "jane.smith@example.com",
        Address: "Bandung, Indonesia",
        Job: "Product Engineer",
    },
    "michael.johnson@example.com": User {
        Email: "michael.johnson@example.com",
        Address: "South Jakarta, Indonesia",
        Job: "QA Engineer",
    },
    "sarah.wilson@example.com": User {
        Email: "sarah.wilson@example.com",
        Address: "Jogjakarta, Indonesia",
        Job: "Security Engineer",
    },
    "robert.jackson@example.com": User {
        Email: "robert.jackson@example.com",
        Address: "er",
        Job: "er",
    },
}
func GetByEmail(email string) (User, error) {
    user, doExist := user[email]

    if(!doExist) {
        return User{}, errors.New("user is not found")
    }

    return user, nil
}

func GetEmails() []string {
    keys := []string {}
    for key := range user {
        keys = append(keys, key)
    }

    return keys
}
