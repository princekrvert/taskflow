package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/princekrvert/taskflow/internal/ui"
)

var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		title := strings.Join(args, " ")
		id, err := taskSvc.AddTask(title)
		if err != nil {
			return err
		}
		fmt.Println(ui.CursorStyle.Render(fmt.Sprintf("✓ Added task #%d: %s", id, title)))
		return nil
	},
}

func init() { rootCmd.AddCommand(addCmd) }
