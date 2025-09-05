# ✅ CLI Todo App

A simple, lightweight **command-line Todo application** written in Go.
It lets you manage tasks (add, edit, delete, toggle) with persistence in a JSON file and a clean tabular output.

---

## 🚀 Features
- **Add** new tasks
- **Edit** existing tasks
- **Toggle** tasks between complete/incomplete
- **Delete** tasks
- **List** all tasks in a pretty table
- **Persistent storage** in a json file

---

## ⚙️ Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/usualboringuy/cli-todo-app.git
   cd cli-todo-app
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Build the executable:
   ```bash
   go build -o todo
   ```
4. Run:
   ```bash
   ./todo -list
   ```

---

## 🛠️ Usage

The app supports the following flags:

| Flag | Description | Example |
|------|-------------|---------|
| `-add "title"` | Add a new task | `./todo -add "Buy groceries"` |
| `-edit "id:new_title"` | Edit a task’s title | `./todo -edit "1:Read DBMS notes"` |
| `-toggle id` | Toggle task status (done/undone) | `./todo -toggle 1` |
| `-del id` | Delete a task by id | `./todo -del 2` |
| `-list` | Show all tasks in a table | `./todo -list` |

---

## 🌍 Cross Compilation

Go makes it easy to build binaries for different platforms without extra tools.
From your project root, run:

### Linux → Windows (`.exe`)
```bash
GOOS=windows GOARCH=amd64 go build -o todo.exe
```

### Linux → Linux (64-bit)
```bash
GOOS=linux GOARCH=amd64 go build -o todo
```

---

### 📌 Notes
- Replace `todo` with any name you prefer.
- On Windows, you’ll usually want the `.exe` extension.

---

## 📜 License
MIT License © 2025 Anish Pechetti
