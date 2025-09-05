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
You can install `todo` **globally in your local user folder** (no admin rights required).

### Linux / macOS
1. Create a local bin folder:
   ```bash
   mkdir -p ~/.local/bin
   ```
2. Build the binary there:
   ```bash
   go build -o ~/.local/bin/todo ./src
   ```
3. Add it to your PATH (`~/.bashrc`, `~/.zshrc`, etc.):
   ```bash
   echo 'export PATH=$HOME/.local/bin:$PATH' >> ~/.bashrc
   source ~/.bashrc
   ```
   or

   ```bash
   echo 'export PATH=$HOME/.local/bin:$PATH' >> ~/.zshrc
   source ~/.zshrc
   ```
Now you can run:
```bash
todo -list
```

---

### Windows
1. Create a local bin folder:
   ```cmd
   mkdir %USERPROFILE%\bin
   ```
2. Build the binary there:
   ```cmd
   go build -o %USERPROFILE%\bin\todo.exe ./src
   ```
3. Add it to PATH (permanent):
   ```cmd
   setx PATH "%PATH%;%USERPROFILE%\bin"
   ```
4. Restart PowerShell / CMD.

Now you can run:
```cmd
todo -list
```

---

## 🛠️ Usage

The app supports the following flags:

| Flag | Description | Example |
|------|-------------|---------|
| `-add "title"` | Add a new task | `todo -add "Buy groceries"` |
| `-edit "id:new_title"` | Edit a task’s title | `todo -edit "0:Read DBMS notes"` |
| `-toggle id` | Toggle task status (done/undone) | `todo -toggle 0` |
| `-del id` | Delete a task by id | `todo -del 0` |
| `-list` | Show all tasks in a table | `todo -list` |

---

### 📌 Notes
- Replace `todo` with any name you prefer.
- On Windows, you’ll usually want the `.exe` extension.
