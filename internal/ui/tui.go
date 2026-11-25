package ui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/princekrvert/taskflow/internal/models"
	"github.com/princekrvert/taskflow/internal/service"
)

var (
	// Palette inspired by Charm's lipgloss-example
	highlight = lipgloss.Color("#7D56F4")
	hotPink   = lipgloss.Color("#F25D94")
	green     = lipgloss.Color("#43BF6D")
	subtle    = lipgloss.Color("#626262")
	barBg     = lipgloss.Color("#353533")

	// Tabs
	activeTabBorder = lipgloss.Border{
		Top: "─", Bottom: " ", Left: "│", Right: "│",
		TopLeft: "╭", TopRight: "╮", BottomLeft: "┘", BottomRight: "└",
	}
	tabBorder = lipgloss.Border{
		Top: "─", Bottom: "─", Left: "│", Right: "│",
		TopLeft: "╭", TopRight: "╮", BottomLeft: "┴", BottomRight: "┴",
	}
	activeTabStyle = lipgloss.NewStyle().Border(activeTabBorder, true).BorderForeground(highlight).Foreground(lipgloss.Color("#FFFDF5")).Bold(true).Padding(0, 2)
	tabStyle       = lipgloss.NewStyle().Border(tabBorder, true).BorderForeground(subtle).Foreground(subtle).Padding(0, 2)

	// List styles
	checkMark   = lipgloss.NewStyle().Foreground(green).SetString("✓")
	doneText    = lipgloss.NewStyle().Strikethrough(true).Foreground(subtle)
	cursorText  = lipgloss.NewStyle().Foreground(hotPink).Bold(true)
	dateTag     = lipgloss.NewStyle().Foreground(subtle)
	sectionHead = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderBottom(true).BorderForeground(subtle).Bold(true).MarginBottom(1)

	// Status bar (bottom)
	statusNugget = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFDF5")).Padding(0, 1).Bold(true)
	statusKey    = statusNugget.Background(hotPink)
	statusMode   = statusNugget.Background(highlight)
	statusVal    = lipgloss.NewStyle().Foreground(lipgloss.Color("#C1C6B2")).Background(barBg).Padding(0, 1)

	// Exported for CLI commands (add.go, stats.go)
	HeaderStyle = lipgloss.NewStyle().Bold(true).Foreground(highlight).MarginBottom(1)
	CursorStyle = lipgloss.NewStyle().Foreground(green).Bold(true)
)

type model struct {
	svc       *service.TaskService
	activeTab int // 0 = Today, 1 = History, 2 = Stats
	today     []models.Task
	history   []models.Task
	stats     []models.MonthlyStat
	cursor    int
	statusMsg string
}

func InitialModel(svc *service.TaskService) model {
	m := model{svc: svc, statusMsg: "Ready"}
	m.refresh()
	return m
}

func (m *model) refresh() {
	todayStr := time.Now().Format("2006-01-02")
	m.today, _ = m.svc.GetDailyTasks(todayStr)
	m.history, _ = m.svc.GetPreviousTasks()
	m.stats, _ = m.svc.MonthlyReport("")
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab", "right", "l":
			m.activeTab = (m.activeTab + 1) % 3
			m.cursor = 0
			m.statusMsg = "Switched tab"
		case "shift+tab", "left", "h":
			m.activeTab = (m.activeTab + 2) % 3
			m.cursor = 0
			m.statusMsg = "Switched tab"
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			maxIdx := m.currentListLen() - 1
			if m.cursor < maxIdx {
				m.cursor++
			}
		case " ", "enter":
			if m.activeTab == 0 && len(m.today) > 0 {
				t := m.today[m.cursor]
				if err := m.svc.CompleteTask(t.ID); err != nil {
					m.statusMsg = err.Error()
				} else {
					m.statusMsg = fmt.Sprintf("Completed #%d", t.ID)
					m.refresh()
				}
			} else if m.activeTab == 1 {
				m.statusMsg = "Read-only: previous tasks cannot be modified"
			}
		case "d":
			if m.activeTab == 0 && len(m.today) > 0 {
				t := m.today[m.cursor]
				if err := m.svc.DeleteTask(t.ID); err != nil {
					m.statusMsg = err.Error()
				} else {
					m.statusMsg = fmt.Sprintf("Deleted #%d", t.ID)
					m.refresh()
					if m.cursor >= len(m.today) && m.cursor > 0 {
						m.cursor--
					}
				}
			} else if m.activeTab == 1 {
				m.statusMsg = "Read-only: previous tasks cannot be deleted"
			}
		}
	}
	return m, nil
}

func (m model) currentListLen() int {
	switch m.activeTab {
	case 0:
		return len(m.today)
	case 1:
		return len(m.history)
	default:
		return len(m.stats)
	}
}

func (m model) View() string {
	var b strings.Builder

	// 1. Tab Bar
	tabs := []string{"Today", "History", "Monthly Stats"}
	var renderedTabs []string
	for i, name := range tabs {
		if i == m.activeTab {
			renderedTabs = append(renderedTabs, activeTabStyle.Render(name))
		} else {
			renderedTabs = append(renderedTabs, tabStyle.Render(name))
		}
	}
	b.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...) + "\n\n")

	// 2. Active Tab Content
	switch m.activeTab {
	case 0:
		b.WriteString(sectionHead.Render("Today's Tasks ("+time.Now().Format("2006-01-02")+")") + "\n")
		if len(m.today) == 0 {
			b.WriteString(dateTag.Render("  No tasks added today. Run `taskflow add \"Your task\"`.\n"))
		}
		for i, t := range m.today {
			b.WriteString(renderTaskRow(t, i == m.cursor, false))
		}
	case 1:
		b.WriteString(sectionHead.Render("Previous Tasks (Read-Only)") + "\n")
		if len(m.history) == 0 {
			b.WriteString(dateTag.Render("  No previous tasks found.\n"))
		}
		for i, t := range m.history {
			b.WriteString(renderTaskRow(t, i == m.cursor, true))
		}
	case 2:
		b.WriteString(sectionHead.Render("Monthly Completion Report") + "\n")
		if len(m.stats) == 0 {
			b.WriteString(dateTag.Render("  No completed tasks recorded yet.\n"))
		}
		for _, s := range m.stats {
			b.WriteString(fmt.Sprintf("  %s  %s tasks completed\n", cursorText.Render(s.Month), checkMark.Render(fmt.Sprintf("%d", s.CompletedCount))))
		}
	}

	// 3. Bottom Status Bar
	modeLabel := "EDITABLE"
	if m.activeTab != 0 {
		modeLabel = "READ-ONLY"
	}
	statusBar := lipgloss.JoinHorizontal(
		lipgloss.Top,
		statusKey.Render("STATUS"),
		statusVal.Width(42).Render(m.statusMsg),
		statusMode.Render(modeLabel),
	)

	b.WriteString("\n" + statusBar + "\n")
	b.WriteString(dateTag.Render("←/→ or tab: switch view • ↑/↓: move • space: complete • d: delete • q: quit"))
	return b.String()
}

func renderTaskRow(t models.Task, selected bool, showDate bool) string {
	cursor := "  "
	if selected {
		cursor = cursorText.Render("▸ ")
	}
	prefix := "  "
	title := t.Title
	if t.Completed {
		prefix = checkMark.String() + " "
		title = doneText.Render(t.Title)
	}
	extra := ""
	if showDate {
		extra = dateTag.Render(" [" + t.CreatedAt.Local().Format("2006-01-02") + "]")
	}
	return fmt.Sprintf("%s%s#%d %s%s\n", cursor, prefix, t.ID, title, extra)
}