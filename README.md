TaskFlowA lightweight, fast CLI task manager built with Go, Cobra, SQLite, Bubble Tea, and Lipgloss. Manage your daily tasks directly from the terminal with a Dracula-inspired UI and track your monthly completion metrics—on desktop or Android (Termux).FeaturesFast CLI Commands: Quickly add, complete (com), delete (del), and filter tasks by day (today).Interactive TUI Dashboard: Run taskflow with no arguments to launch a keyboard-driven Bubble Tea interface.Daily Styled View: View today's (or any past date's) tasks inside a rounded Lipgloss card with color-coded TODO/DONE badges and a live progress counter.Monthly Productivity Analytics: Automatically timestamps completed tasks in SQLite so you can review your monthly output anytime.Zero-Config & Cross-Platform: Powered by pure-Go SQLite (modernc.org/sqlite) with no CGO or external database setup required. Works out of the box on Linux, macOS, Windows, and Android (Termux).Installation GuidePrerequisitesGo 1.22 or higher1. Quick Install via Go (Linux / macOS / Windows)go install github.com/princekrvert/taskflow@latest
(Ensure your $GOPATH/bin or ~/go/bin is added to your system PATH.)2. Build from SourceLinux & macOS:git clone https://github.com/princekrvert/taskflow.git
cd taskflow
go mod tidy
go build -o taskflow .
sudo mv taskflow /usr/local/bin/
Windows (PowerShell):git clone https://github.com/princekrvert/taskflow.git
cd taskflow
go mod tidy
go build -o taskflow.exe .
Move taskflow.exe to a folder included in your Windows PATH.3. Install on Android (Termux)Option A: Build Directly on Termux (Recommended)pkg update && pkg install golang git -y
git clone https://github.com/princekrvert/taskflow.git
cd taskflow
go mod tidy
go build -o $PREFIX/bin/taskflow .
Option B: Cross-Compile from PCBuild the Android ARM64 binary on your computer:GOOS=android GOARCH=arm64 go build -o taskflow .
Transfer the taskflow binary to your phone, then move it into Termux's PATH:mv /sdcard/Download/taskflow $PREFIX/bin/
chmod +x $PREFIX/bin/taskflow
Command Reference & UsageInteractive TUI ModeLaunch the interactive Bubble Tea dashboard by running taskflow without any subcommands:taskflow
KeyAction↑ / kMove cursor up↓ / jMove cursor downSpace / EnterMark selected task as completeddDelete selected taskq / Ctrl+CQuit dashboardCLI Commands OverviewCommandAliasesFlagsDescriptiontaskflow add <title>——Add a new tasktaskflow todaydaily, ls-d, --dateView tasks for today or a specific date (YYYY-MM-DD)taskflow com <id>done, complete—Mark a task as completed by its IDtaskflow del <id>rm, delete—Permanently delete a task by its IDtaskflow stats—-m, --monthView monthly completed task counts (YYYY-MM)Usage ExamplesAdd a new task:taskflow add "Finish Go CLI project"
View today's tasks (styled Lipgloss view):taskflow today
# or using aliases:
taskflow daily
taskflow ls
View tasks for a specific date:taskflow today -d 2026-11-20
Mark a task as completed:taskflow com 1
# or:
taskflow done 1
Delete a task:taskflow del 1
# or:
taskflow rm 1
View monthly completion statistics:# All months
taskflow stats

# Specific month
taskflow stats -m 2026-11
Data StorageTaskFlow automatically creates and manages a local SQLite database in your home directory:Linux / macOS / Windows: ~/.taskflow.dbAndroid (Termux): /data/data/com.termux/files/home/.taskflow.dbProject ArchitectureTaskFlow uses a decoupled, layered architecture so business logic, storage, and UI can be modified independently:taskflow/
├── cmd/                   # Cobra CLI commands and flag parsing
├── internal/
│   ├── models/            # Core Task and MonthlyStat structs
│   ├── repository/        # SQLite queries behind the TaskRepository interface
│   ├── service/           # Business logic and input validation
│   └── ui/                # Bubble Tea TUI and Lipgloss daily view styles
├── go.mod
├── main.go
└── README.md
LicenseMIT