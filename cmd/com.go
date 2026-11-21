package cmd

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var comStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#50FA7B")). // Mint Green
	Bold(true)

var comCmd = &cobra.Command{
	Use:     "com [task ID]",
	Aliases: []string{"done", "complete"},
	Short:   "Mark a task as completed by its ID",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid task ID %q: must be a number", args[0])
		}

		if err := taskSvc.CompleteTask(id); err != nil {
			return err
		}

		fmt.Println(comStyle.Render(fmt.Sprintf("✓ Completed task #%d", id)))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(comCmd)
}
