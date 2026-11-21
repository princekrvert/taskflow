package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/princekrvert/taskflow/internal/models"
	"github.com/princekrvert/taskflow/internal/service"
)

var (
	HeaderStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7D56F4")).MarginBottom(1)
	DoneStyle   = lipgloss.NewStyle().Strikethrough(true).Foreground(lipgloss.Color("#626262"))
	CursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575")).Bold(true)
	HelpStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).MarginTop(1)
)

type model struct {
	svc    *service.TaskService
	tasks  []models.Task
	cursor int
	err    error
}

func InitialModel(svc *service.TaskService) model {
	tasks, err := svc.GetTasks()
	return model{svc: svc, tasks: tasks, err: err}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}
		case " ", "enter":
			if len(m.tasks) > 0 {
				_ = m.svc.CompleteTask(m.tasks[m.cursor].ID)
				m.tasks, _ = m.svc.GetTasks()
			}
		case "d":
			if len(m.tasks) > 0 {
				_ = m.svc.DeleteTask(m.tasks[m.cursor].ID)
				m.tasks, _ = m.svc.GetTasks()
				if m.cursor >= len(m.tasks) && m.cursor > 0 {
					m.cursor--
				}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := HeaderStyle.Render("TaskFlow — Interactive Task Manager") + "\n"
	if len(m.tasks) == 0 {
		s += "No tasks yet. Run `taskflow add \"Your task\"` to get started.\n"
	}

	for i, t := range m.tasks {
		cursor := "  "
		if m.cursor == i {
			cursor = CursorStyle.Render("> ")
		}

		checkbox := "[ ]"
		title := t.Title
		if t.Completed {
			checkbox = "[✓]"
			title = DoneStyle.Render(t.Title)
		}

		s += fmt.Sprintf("%s%s #%d %s\n", cursor, checkbox, t.ID, title)
	}

	s += HelpStyle.Render("↑/↓: navigate • space: complete • d: delete • q: quit")
	return s
}
