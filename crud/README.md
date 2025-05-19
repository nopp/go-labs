# Simple CRUD in Go (SQLite)

This project is a minimal CRUD (Create, Read, Update, Delete) application written in Go. It uses:

- **SQLite** as the local database
- **Go's net/http** for web routing
- **HTML templates** with Bootstrap 5 for a clean UI

---

## 📦 Features

- List all products
- Add new product
- Edit product
- Delete product
- Data persisted with SQLite (`products.db`)

---

## 🏗️ Project Structure

```
simple-crud/
├── main.go             # Application entrypoint
├── database.go         # Database initialization and connection
├── handlers.go         # All HTTP handlers (CRUD logic)
├── models.go           # Product struct definition
├── templates/          # HTML templates
│   ├── index.html      # Product listing page
│   └── form.html       # Create/Edit product form
├── Makefile            # Shortcut to run the project
```

---

## 🚀 How to Run

### 1. Install dependencies

```bash
go mod init simple-crud
go get github.com/mattn/go-sqlite3
```

### 2. Run the server

```bash
make run
```

Or directly:

```bash
go run .
```

### 3. Open in browser

Visit: [http://localhost:8080](http://localhost:8080)

---

## 📝 Example Product Fields

- **Name**: Product name (e.g. "Mouse")
- **Description**: Short description (e.g. "Wireless USB mouse")
- **Price**: Decimal price (e.g. 99.90)

---

## 🛠️ Requirements

- Go 1.18+
- SQLite3 (automatically handled by driver)

---

## 📃 License

MIT — feel free to use, modify, and share.