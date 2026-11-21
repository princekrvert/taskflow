package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/princekrvert/taskflow/internal/repository"
	"github.com/princekrvert/taskflow/internal/service"
	"github.com/princekrvert/taskflow/internal/ui"
)

var taskSvc *service.TaskService

var rootCmd = &cobra.Command{
	Use:   "taskflow",
	Short: "A fast, SQLite-backed CLI task manager with monthly tracking",
	RunE: func(cmd *cobra.Command, args []string) error {
		p := tea.NewProgram(ui.InitialModel(taskSvc))
		_, err := p.Run()
		return err
	},
}

func Execute() {
	home, _ := os.UserHomeDir()
	dbPath := filepath.Join(home, ".taskflow.db")

	repo, err := repository.NewSQLiteRepo(dbPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize database: %v\n", err)
		os.Exit(1)
	}
	taskSvc = service.NewTaskService(repo)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
