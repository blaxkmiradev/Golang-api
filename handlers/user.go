package handlers

import (
    "encoding/json"
    "io/ioutil"
    "myapi/models"
    "net/http"
    "strconv"
    "strings"
)

var users = []models.User{
    {ID: 1, Name: "Alice", Age: 25},
    {ID: 2, Name: "Bob", Age: 30},
}

// Helper: find user index by ID
func findUserIndex(id int) int {
    for i, u := range users {
        if u.ID == id {
            return i
        }
    }
    return -1
}

// GET /users  or POST /users
func UsersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    switch r.Method {
    case http.MethodGet:
        json.NewEncoder(w).Encode(users)
    case http.MethodPost:
        var newUser models.User
        body, _ := ioutil.ReadAll(r.Body)
        json.Unmarshal(body, &newUser)

        newUser.ID = len(users) + 1
        users = append(users, newUser)
        json.NewEncoder(w).Encode(newUser)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// GET /users/{id}, PUT /users/{id}, DELETE /users/{id}
func UserHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    idStr := strings.TrimPrefix(r.URL.Path, "/users/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    index := findUserIndex(id)
    if index == -1 {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    switch r.Method {
    case http.MethodGet:
        json.NewEncoder(w).Encode(users[index])
    case http.MethodPut:
        var updatedUser models.User
        body, _ := ioutil.ReadAll(r.Body)
        json.Unmarshal(body, &updatedUser)
        updatedUser.ID = id
        users[index] = updatedUser
        json.NewEncoder(w).Encode(updatedUser)
    case http.MethodDelete:
        users = append(users[:index], users[index+1:]...)
        json.NewEncoder(w).Encode(map[string]string{"message": "User deleted"})
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
