🚀 MyAPI — Simple Go CRUD API

A lightweight REST API built in Go using only the standard library. Perfect for learning, testing, or small projects.

✅ CRUD operations for users
✅ No external dependencies
✅ Easy to extend to persistent storage
📂 Project Structure
myapi/
├── main.go          # Entry point
├── go.mod           # Go module
├── handlers/        # HTTP handlers
│   └── user.go
├── models/          # Structs for data
│   └── user.go
└── data/            # Optional: JSON file for persistence
    └── users.json
⚙️ Setup & Run
Clone the repo (or create your folder):
git clone https://github.com/blaxkmiradev/Golang-api.git
cd Golang-api
Initialize Go module (if not done):
go mod init myapi
Run the API server:
go run main.go

Server runs on http://localhost:8080

🌐 API Endpoints
Get all users
GET /users

Response:

[
  {"id":1,"name":"Alice","age":25},
  {"id":2,"name":"Bob","age":30}
]
Get user by ID
GET /users/{id}

Example:

curl http://localhost:8080/users/1

Response:

{"id":1,"name":"Alice","age":25}
Create a new user
POST /users
Content-Type: application/json

Body:

{
  "name": "Charlie",
  "age": 28
}

Response:

{
  "id": 3,
  "name": "Charlie",
  "age": 28
}
Update an existing user
PUT /users/{id}
Content-Type: application/json

Body:

{
  "name": "Alice Updated",
  "age": 26
}

Response:

{
  "id": 1,
  "name": "Alice Updated",
  "age": 26
}
Delete a user
DELETE /users/{id}

Response:

{"message": "User deleted"}
