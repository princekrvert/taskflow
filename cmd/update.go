package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/princekrvert/taskflow/internal/ui"
	"github.com/spf13/cobra"
)

type githubRelease struct {
	TagName string `json:"tag_name"`
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check GitHub for a newer version of TaskFlow and update",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Current version: %s\nChecking GitHub for updates...\n", Version)

		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Get("https://api.github.com/repos/princekrvert/taskflow/releases/latest")
		if err != nil {
			return fmt.Errorf("failed to reach GitHub: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound {
			fmt.Println("No GitHub releases published yet. Pulling latest main branch...")
		} else if resp.StatusCode == http.StatusOK {
			var rel githubRelease
			if err := json.NewDecoder(resp.Body).Decode(&rel); err == nil {
				if rel.TagName == Version {
					fmt.Println(ui.CursorStyle.Render("✓ TaskFlow is already up to date (" + Version + ")"))
					return nil
				}
				fmt.Printf("New version available: %s → updating...\n", rel.TagName)
			}
		}

		install := exec.Command("go", "install", "github.com/princekrvert/taskflow@latest")
		install.Stdout = os.Stdout
		install.Stderr = os.Stderr
		if err := install.Run(); err != nil {
			return fmt.Errorf("update failed: %w", err)
		}

		fmt.Println(ui.CursorStyle.Render("✓ TaskFlow updated successfully!"))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
