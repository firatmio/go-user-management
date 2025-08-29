package main

import (
    "encoding/json"
    "net/http"
    "golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var u User
    json.NewDecoder(r.Body).Decode(&u)

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

    _, err := DB.Exec("INSERT INTO users(username, password) VALUES(?, ?)", u.Username, string(hashedPassword))
    if err != nil {
        http.Error(w, "User already exists", http.StatusConflict)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("User registered successfully"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var u User
    json.NewDecoder(r.Body).Decode(&u)

    row := DB.QueryRow("SELECT password FROM users WHERE username = ?", u.Username)
    var hashedPassword string
    err := row.Scan(&hashedPassword)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(u.Password))
    if err != nil {
        http.Error(w, "Invalid password", http.StatusUnauthorized)
        return
    }

    w.Write([]byte("Login successful"))
}
