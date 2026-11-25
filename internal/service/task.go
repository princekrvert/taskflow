package service

import (
	"errors"
	"strings"
	"time"

	"github.com/princekrvert/taskflow/internal/models"
	"github.com/princekrvert/taskflow/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) isToday(t time.Time) bool {
	return t.Local().Format("2006-01-02") == time.Now().Format("2006-01-02")
}

func (s *TaskService) AddTask(title string) (int64, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return 0, errors.New("task title cannot be empty")
	}
	return s.repo.Create(title)
}

func (s *TaskService) DeleteTask(id int) error {
	if id <= 0 {
		return errors.New("invalid task ID")
	}
	task, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("task not found")
	}
	if !s.isToday(task.CreatedAt) {
		return errors.New("locked: cannot delete tasks from previous days")
	}
	return s.repo.Delete(id)
}

func (s *TaskService) CompleteTask(id int) error {
	if id <= 0 {
		return errors.New("invalid task ID")
	}
	task, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("task not found")
	}
	if !s.isToday(task.CreatedAt) {
		return errors.New("locked: cannot modify tasks from previous days")
	}
	return s.repo.MarkComplete(id)
}

func (s *TaskService) GetTasks() ([]models.Task, error) {
	return s.repo.ListAll()
}

func (s *TaskService) GetDailyTasks(date string) ([]models.Task, error) {
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}
	return s.repo.ListByDate(date)
}

func (s *TaskService) GetPreviousTasks() ([]models.Task, error) {
	today := time.Now().Format("2006-01-02")
	return s.repo.ListPrevious(today)
}

func (s *TaskService) MonthlyReport(month string) ([]models.MonthlyStat, error) {
	return s.repo.GetMonthlyStats(month)
}
