package models

import "time"

type Task struct {
	ID          int
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time // Pointer allows NULL in SQLite
}

type MonthlyStat struct {
	Month          string // Format: YYYY-MM
	CompletedCount int
}
