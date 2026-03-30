MyAPI — Simple Go API

A small REST API in Go for managing users.

Run the API
Make sure you have Go installed (1.20+ recommended).
Open a terminal in the folder where main.go is located.
Run the server:
go run main.go

The server will start on:

http://localhost:8080
Test the API
Get all users:
curl http://localhost:8080/users
Get user by ID:
curl http://localhost:8080/users/1
Add a new user:
curl -X POST -d '{"name":"Charlie","age":28}' -H "Content-Type: application/json" http://localhost:8080/users
Update user:
curl -X PUT -d '{"name":"Alice Updated","age":26}' -H "Content-Type: application/json" http://localhost:8080/users/1
Delete user:
curl -X DELETE http://localhost:8080/users/1
