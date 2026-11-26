# TaskFlow

<p align="center">
  <strong>A simple, keyboard-driven task manager for your terminal.</strong>
  <br />
  Built with Go, Bubble Tea, Cobra, Lipgloss, and SQLite.
</p>

<p align="center">
  <a href="https://github.com/princekrvert/taskflow/stargazers">
    <img src="https://img.shields.io/github/stars/princekrvert/taskflow?style=flat-square" alt="GitHub stars" />
  </a>
  <a href="https://github.com/princekrvert/taskflow/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/princekrvert/taskflow?style=flat-square" alt="License" />
  </a>
  <img src="https://img.shields.io/badge/Go-1.26.0-00ADD8?style=flat-square&logo=go" alt="Go version" />
  <img src="https://img.shields.io/badge/Database-SQLite-003B57?style=flat-square&logo=sqlite" alt="SQLite" />
</p>

---

## Overview

**TaskFlow** is a lightweight command-line and terminal UI task manager designed to help you organize your daily work without leaving your terminal.

Whether you prefer executing commands or navigating an interactive terminal dashboard, TaskFlow provides a straightforward way to manage tasks and track your productivity.

Your tasks are stored locally in an SQLite database, so there's no need for an external database server or cloud account.

## Features

- **Command-Line Interface:** Add, complete, delete, and inspect tasks directly from your terminal.
- **Interactive TUI:** Launch a full-screen task dashboard with Bubble Tea.
- **Keyboard Navigation:** Navigate tasks using Vim-style `j` and `k` shortcuts.
- **Daily Task Tracking:** View tasks for today or inspect a specific date.
- **Completion Tracking:** Automatically record task completion timestamps.
- **Monthly Statistics:** Review how many tasks you've completed each month.
- **Local Persistence:** Store tasks in an SQLite database on your machine.
- **Zero Database Setup:** Uses a pure-Go SQLite driver without requiring CGO or an external database server.

## Preview

```text
╭────────────────────────────────────────────────╮
│  Daily Tasks                                    │
│                                                │
│  TODO  #4   Finish Go CLI documentation        │
│  TODO  #3   Push TaskFlow to GitHub            │
│  DONE  #2   Build Lipgloss daily view          │
│  DONE  #1   Set up SQLite repository           │
│                                                │
│  Progress: 2/4 completed                       │
╰────────────────────────────────────────────────╯
```

*Illustrative terminal preview.*

## Installation

### Prerequisites

- Go 1.26.0 or later.
- A supported terminal.

### Option 1: Install using Go

```bash
go install github.com/princekrvert/taskflow@latest
```

Make sure your Go binary directory is included in your system's `PATH`.

Run TaskFlow:

```bash
taskflow
```

### Option 2: Build from source

Clone the repository:

```bash
git clone https://github.com/princekrvert/taskflow.git
cd taskflow
```

Build the application:

```bash
go build -o taskflow .
```

Run it:

```bash
./taskflow
```

To install it system-wide on Linux or macOS, move the binary into a directory in your `PATH`:

```bash
sudo mv taskflow /usr/local/bin/
```

### Option 3: Windows

```powershell
git clone https://github.com/princekrvert/taskflow.git
cd taskflow

go build -o taskflow.exe .
```

Run:

```powershell
.\taskflow.exe
```

### Option 4: Android (Termux)

Install the prerequisites:

```bash
pkg update
pkg install golang git -y
```

Clone and build:

```bash
git clone https://github.com/princekrvert/taskflow.git
cd taskflow

go build -o $PREFIX/bin/taskflow .
```

Run:

```bash
taskflow
```

## Usage

TaskFlow provides two ways to interact with your tasks.

### 1. Interactive TUI

Launch TaskFlow without any arguments:

```bash
taskflow
```

The interactive dashboard lets you navigate and manage your tasks using your keyboard.

| Key | Action |
|---|---|
| `↑` / `k` | Move cursor up |
| `↓` / `j` | Move cursor down |
| `Space` / `Enter` | Mark selected task as completed |
| `d` | Delete selected task |
| `q` / `Ctrl+C` | Quit |

### 2. Command-Line Interface

#### Add a task

Create a new task:

```bash
taskflow add "Learn Go concurrency"
```

#### View daily tasks

Display today's tasks:

```bash
taskflow today
```

Aliases:

```bash
taskflow daily
taskflow ls
```

View tasks for a specific date:

```bash
taskflow today --date 2026-11-20
```

Short flag:

```bash
taskflow today -d 2026-11-20
```

#### Complete a task

Mark a task as completed using its ID:

```bash
taskflow com 1
```

Aliases:

```bash
taskflow done 1
taskflow complete 1
```

#### Delete a task

Delete a task using its ID:

```bash
taskflow del 1
```

Aliases:

```bash
taskflow rm 1
taskflow delete 1
```

#### View productivity statistics

Display monthly task completion statistics:

```bash
taskflow stats
```

Filter statistics by month:

```bash
taskflow stats --month 2026-11
```

Short flag:

```bash
taskflow stats -m 2026-11
```

## Command Reference

| Command | Description |
|---|---|
| `taskflow` | Launch interactive dashboard |
| `taskflow add <description>` | Create a task |
| `taskflow today` | View today's tasks |
| `taskflow today -d <date>` | View tasks for a specific date |
| `taskflow com <id>` | Complete a task |
| `taskflow del <id>` | Delete a task |
| `taskflow stats` | View monthly completion statistics |
| `taskflow stats -m <month>` | View statistics for a specific month |

## Task Management Rules

TaskFlow follows a simple daily task management approach:

- Tasks are created with an automatically generated ID.
- Tasks can be completed and their completion timestamps are recorded.
- Tasks created on previous days are protected from modification through the service layer.
- Daily views can include tasks created on the selected date or completed on that date.
- Monthly statistics are calculated using task completion timestamps.

This keeps historical task records protected while allowing you to review previous productivity.

## Data Storage

TaskFlow stores all task data locally using SQLite.

**Default database location:**

```text
~/.taskflow.db
```

The location is based on your operating system's home directory.

| Platform | Database location |
|---|---|
| Linux | `$HOME/.taskflow.db` |
| macOS | `$HOME/.taskflow.db` |
| Windows | `%USERPROFILE%\.taskflow.db` |
| Android (Termux) | `$HOME/.taskflow.db` |

No external database server is required.

## Architecture

TaskFlow uses a layered architecture that separates command handling, business logic, data access, and terminal presentation.

```text
taskflow/
│
├── cmd/
│   ├── root.go
│   ├── add.go
│   ├── today.go
│   ├── com.go
│   ├── del.go
│   └── stats.go
│
├── internal/
│   ├── models/
│   ├── repository/
│   ├── service/
│   └── ui/
│
├── main.go
├── go.mod
├── go.sum
└── README.md
```

### Architecture responsibilities

| Layer | Responsibility |
|---|---|
| `cmd` | CLI commands, flags, and application entry points |
| `models` | Core task and statistics data structures |
| `repository` | Database operations and repository abstraction |
| `service` | Business rules, validation, and task management |
| `ui` | Interactive terminal interface and visual presentation |
| `main.go` | Application entry point |

### Technology Stack

| Technology | Purpose |
|---|---|
| [Go](https://go.dev/) | Programming language |
| [Cobra](https://github.com/spf13/cobra) | CLI command management |
| [Bubble Tea](https://github.com/charmbracelet/bubbletea) | Interactive terminal interface |
| [Lipgloss](https://github.com/charmbracelet/lipgloss) | Terminal styling |
| [SQLite](https://sqlite.org/) | Local database |
| [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite) | Pure-Go SQLite driver |

## Development

Want to contribute or explore the code?

Clone the project:

```bash
git clone https://github.com/princekrvert/taskflow.git
cd taskflow
```

Download dependencies:

```bash
go mod download
```

Build:

```bash
go build -o taskflow .
```

Run:

```bash
go run .
```

Run tests:

```bash
go test ./...
```

## Contributing

Contributions, suggestions, and bug reports are welcome!

1. Fork the repository.
2. Create a feature branch.
3. Make your changes.
4. Run the tests.
5. Submit a pull request.

```bash
git checkout -b feature/my-feature
```

## Roadmap

Ideas for future improvements:

- [ ] Edit existing tasks.
- [ ] Task priorities and categories.
- [ ] Search and filter tasks.
- [ ] Task deadlines and reminders.
- [ ] Export and import tasks.
- [ ] Improved productivity analytics.
- [ ] Automated tests and CI workflows.

## License

TaskFlow is distributed under the [MIT License](LICENSE).

Created by [@princekrvert](https://github.com/princekrvert).

---

<p align="center">
  Built with Go and a love for the terminal.
  <br />
  <strong>Keep your workflow simple. Stay productive.</strong>
</p>
