package cmd

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var delStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF5555")). // Dracula Red
	Bold(true)

var delCmd = &cobra.Command{
	Use:     "del [task ID]",
	Aliases: []string{"rm", "delete"},
	Short:   "Delete a task by its ID",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid task ID %q: must be a number", args[0])
		}

		if err := taskSvc.DeleteTask(id); err != nil {
			return err
		}

		fmt.Println(delStyle.Render(fmt.Sprintf("✗ Deleted task #%d", id)))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
