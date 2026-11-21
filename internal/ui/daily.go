package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/princekrvert/taskflow/internal/models"
)

var (
	// Container card
	dailyBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")). // Violet
			Padding(1, 2).
			MarginTop(1)

	// Header & Date Pill
	dailyTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#F8F8F2"))

	dateBadgeStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#1E1E2E")).
			Background(lipgloss.Color("#8BE9FD")). // Cyan
			Padding(0, 1).
			MarginLeft(1)

	// Status Badges
	doneBadge = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#1E1E2E")).
			Background(lipgloss.Color("#50FA7B")). // Mint Green
			Padding(0, 1).
			Render("DONE")

	todoBadge = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#1E1E2E")).
			Background(lipgloss.Color("#FFB86C")). // Amber
			Padding(0, 1).
			Render("TODO")

	// Task Text
	idStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#6272A4")).Bold(true)
	pendingTitle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F8F8F2"))
	doneTitle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#6272A4")).Strikethrough(true)
	footerStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#BD93F9")).Italic(true).MarginTop(1)
)

func RenderDailyView(date string, tasks []models.Task) string {
	var b strings.Builder

	header := dailyTitleStyle.Render("Daily Tasks") + dateBadgeStyle.Render(date)
	b.WriteString(header + "\n\n")

	if len(tasks) == 0 {
		b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#6272A4")).Render("No tasks recorded for this day."))
		return dailyBoxStyle.Render(b.String())
	}

	completedCount := 0
	for _, t := range tasks {
		id := idStyle.Render(fmt.Sprintf("#%-3d", t.ID))
		if t.Completed {
			completedCount++
			b.WriteString(fmt.Sprintf("%s %s %s\n", doneBadge, id, doneTitle.Render(t.Title)))
		} else {
			b.WriteString(fmt.Sprintf("%s %s %s\n", todoBadge, id, pendingTitle.Render(t.Title)))
		}
	}

	summary := fmt.Sprintf("Progress: %d/%d completed", completedCount, len(tasks))
	b.WriteString(footerStyle.Render(summary))

	return dailyBoxStyle.Render(b.String())
}
