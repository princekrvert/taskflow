package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/princekrvert/taskflow/internal/ui"
)

var monthFlag string

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "View monthly completed task statistics",
	RunE: func(cmd *cobra.Command, args []string) error {
		stats, err := taskSvc.MonthlyReport(monthFlag)
		if err != nil {
			return err
		}
		fmt.Println(ui.HeaderStyle.Render("Monthly Completion Report"))
		if len(stats) == 0 {
			fmt.Println("No completed tasks found for this period.")
			return nil
		}
		for _, s := range stats {
			fmt.Printf("• %s : %d tasks completed\n", s.Month, s.CompletedCount)
		}
		return nil
	},
}

func init() {
	statsCmd.Flags().StringVarP(&monthFlag, "month", "m", "", "Filter by month (YYYY-MM)")
	rootCmd.AddCommand(statsCmd)
}
