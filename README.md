# TaskFlow

A fast, keyboard-driven CLI and TUI task manager built with **Go**, **Cobra**, **SQLite**, and **Bubble Tea**. Track daily tasks from your terminal and monitor monthly completion metrics on Linux, macOS, Windows, and Android (Termux).

[![Go Version](https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev/)
[![SQLite](https://img.shields.io/badge/Database-SQLite%20(Pure%20Go)-003B57?style=flat-square&logo=sqlite&logoColor=white)](https://pkg.go.dev/modernc.org/sqlite)
[![Bubble Tea](https://img.shields.io/badge/TUI-Bubble%20Tea%20%2B%20Lipgloss-FF79C6?style=flat-square)](https://github.com/charmbracelet/bubbletea)
[![License: MIT](https://img.shields.io/badge/License-MIT-BD93F9?style=flat-square)](#license)

---

## Preview

```text
╭────────────────────────────────────────────────╮
│  Daily Tasks  2026-11-21                       │
│                                                │
│  TODO  #4   Finish Go CLI documentation        │
│  TODO  #3   Push TaskFlow to GitHub            │
│  DONE  #2   Build Lipgloss daily view          │
│  DONE  #1   Set up SQLite repository           │
│                                                │
│  Progress: 2/4 completed                       │
╰────────────────────────────────────────────────╯
```

---

## Features

* **Fast CLI Workflow:** Add, complete, delete, and inspect tasks directly from the command line.
* **Interactive TUI Dashboard:** Launch `taskflow` with no arguments for a full-screen Bubble Tea interface with Vim-style (`j`/`k`) navigation.
* **Daily Styled View:** View today's (or any specific date's) tasks inside a Dracula-themed Lipgloss card with live completion progress.
* **Monthly Productivity Tracking:** Completed tasks are automatically timestamped in SQLite so you can review monthly completion counts.
* **Zero-Config Storage:** Uses pure-Go SQLite (`modernc.org/sqlite`) with no CGO or external database server required.

---

## Installation

### Prerequisites

* Go **1.22** or higher

### Option 1: Install via Go (Linux / macOS / Windows)

```bash
go install [github.com/princekrvert/taskflow@latest](https://github.com/princekrvert/taskflow@latest)
```

### Option 2: Build from Source (Linux / macOS)

```bash
git clone [https://github.com/princekrvert/taskflow.git](https://github.com/princekrvert/taskflow.git)
cd taskflow
go build -o taskflow .
sudo mv taskflow /usr/local/bin/
```

### Option 3: Build from Source (Windows PowerShell)

```powershell
git clone [https://github.com/princekrvert/taskflow.git](https://github.com/princekrvert/taskflow.git)
cd taskflow
go build -o taskflow.exe .
```

Move `taskflow.exe` to a directory included in your system `PATH`.

### Option 4: Install on Android (Termux)

**Build directly on Termux (Recommended):**

```bash
pkg update && pkg install golang git -y
git clone [https://github.com/princekrvert/taskflow.git](https://github.com/princekrvert/taskflow.git)
cd taskflow
go build -o $PREFIX/bin/taskflow .
```

**Or cross-compile from a PC:**

```bash
# On your PC
GOOS=android GOARCH=arm64 go build -o taskflow .

# Copy the binary to your phone, then run in Termux:
mv /sdcard/Download/taskflow $PREFIX/bin/
chmod +x $PREFIX/bin/taskflow
```

---

## Usage

### Interactive TUI Dashboard

Run `taskflow` without any subcommands to open the interactive terminal interface:

```bash
taskflow
```

| Key | Action |
| :--- | :--- |
| `↑` / `k` | Move cursor up |
| `↓` / `j` | Move cursor down |
| `Space` / `Enter` | Mark selected task as completed |
| `d` | Delete selected task |
| `q` / `Ctrl+C` | Quit |

### CLI Commands

| Command | Aliases | Description | Example |
| :--- | :--- | :--- | :--- |
| `add <text>` | — | Add a new task | `taskflow add "Finish Go CLI project"` |
| `today` | `daily`, `ls` | View all tasks for today | `taskflow today` |
| `today -d <YYYY-MM-DD>` | — | View tasks for a specific date | `taskflow today -d 2026-11-20` |
| `com <id>` | `done`, `complete` | Mark a task as completed by ID | `taskflow com 1` |
| `del <id>` | `rm`, `delete` | Delete a task by ID | `taskflow del 1` |
| `stats` | — | Show monthly completed task counts | `taskflow stats` |
| `stats -m <YYYY-MM>` | — | Filter stats for a specific month | `taskflow stats -m 2026-11` |

---

## Architecture

TaskFlow follows a decoupled, layered architecture so the CLI, UI, business logic, and database can be modified or replaced independently:

```text
taskflow/
├── cmd/                   # Cobra CLI command definitions and flag parsing
│   ├── root.go            # Root command (launches Bubble Tea TUI)
│   ├── add.go             # Add task command
│   ├── today.go           # Daily task view command
│   ├── com.go             # Mark task complete command
│   ├── del.go             # Delete task command
│   └── stats.go           # Monthly completion stats command
├── internal/
│   ├── models/            # Core domain structs (Task, MonthlyStat)
│   ├── repository/        # SQLite implementation behind TaskRepository interface
│   ├── service/           # Business logic and input validation
│   └── ui/                # Bubble Tea TUI and Lipgloss styling
├── go.mod
└── main.go                # Application entry point
```

### Data Storage

All tasks are stored locally in a SQLite database at:

* **Linux / macOS:** `~/.taskflow.db`
* **Windows:** `%USERPROFILE%\.taskflow.db`
* **Termux:** `$HOME/.taskflow.db`

---

## License

Distributed under the MIT License. Created by [@princekrvert](https://github.com/princekrvert).