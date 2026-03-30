package main

import (
    "fmt"
    "myapi/handlers"
    "net/http"
)

func main() {
    http.HandleFunc("/users", handlers.UsersHandler)
    http.HandleFunc("/users/", handlers.UserHandler) // for /users/{id}

    fmt.Println("Server running on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}
