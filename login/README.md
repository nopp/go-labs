# Simple Login MVC

This is a sample project written in Go with SQLite using the **MVC** (Model-View-Controller) architectural pattern.

## 📦 Project Structure

```
simple-login/
├── main.go                 # Entry point
├── go.mod
├── db/                    # Database connection and initialization
│   └── database.go
├── models/                # Business logic and database access
│   └── user.go
├── controllers/           # HTTP request handlers
│   ├── auth.go
│   └── home.go
├── views/                 # HTML templates
│   ├── login.html
│   └── home.html
```

## ✅ Application Flow

- `main.go` sets up the routes and starts the HTTP server
- `controllers/` handles incoming HTTP requests and controls flow
- `models/` provides business logic and database access
- `views/` contains HTML templates rendered as responses

## 🔐 Default Login

- **Username:** admin
- **Password:** 1234

## ▶️ Run the Project

```bash
go mod tidy
go run main.go
```

Access: [http://localhost:8080](http://localhost:8080)