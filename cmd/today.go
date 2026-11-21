package cmd

import (
	"fmt"
	"time"

	"github.com/princekrvert/taskflow/internal/ui"
	"github.com/spf13/cobra"
)

var dateFlag string

var todayCmd = &cobra.Command{
	Use:     "today",
	Aliases: []string{"daily", "ls"},
	Short:   "View all tasks for today (or a specific date)",
	RunE: func(cmd *cobra.Command, args []string) error {
		targetDate := dateFlag
		if targetDate == "" {
			targetDate = time.Now().Format("2006-01-02")
		}

		tasks, err := taskSvc.GetDailyTasks(targetDate)
		if err != nil {
			return err
		}

		fmt.Println(ui.RenderDailyView(targetDate, tasks))
		return nil
	},
}

func init() {
	todayCmd.Flags().StringVarP(&dateFlag, "date", "d", "", "Filter by specific date (YYYY-MM-DD)")
	rootCmd.AddCommand(todayCmd)
}
