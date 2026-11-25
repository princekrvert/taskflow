package repository

import (
	"database/sql"
	"time"

	"github.com/princekrvert/taskflow/internal/models"
	_ "modernc.org/sqlite"
)

type TaskRepository interface {
	Create(title string) (int64, error)
	Delete(id int) error
	MarkComplete(id int) error
	GetByID(id int) (*models.Task, error)
	ListAll() ([]models.Task, error)
	ListByDate(date string) ([]models.Task, error)
	ListPrevious(today string) ([]models.Task, error)
	GetMonthlyStats(month string) ([]models.MonthlyStat, error)
}

type SQLiteRepo struct {
	db *sql.DB
}

func NewSQLiteRepo(dbPath string) (*SQLiteRepo, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		completed BOOLEAN DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		completed_at DATETIME
	);`

	if _, err := db.Exec(query); err != nil {
		return nil, err
	}

	return &SQLiteRepo{db: db}, nil
}

func (r *SQLiteRepo) Create(title string) (int64, error) {
	res, err := r.db.Exec(`INSERT INTO tasks (title) VALUES (?)`, title)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *SQLiteRepo) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM tasks WHERE id = ?`, id)
	return err
}

func (r *SQLiteRepo) MarkComplete(id int) error {
	_, err := r.db.Exec(
		`UPDATE tasks SET completed = 1, completed_at = ? WHERE id = ?`,
		time.Now().UTC(), id,
	)
	return err
}

func (r *SQLiteRepo) ListAll() ([]models.Task, error) {
	rows, err := r.db.Query(`SELECT id, title, completed, created_at, completed_at FROM tasks ORDER BY completed ASC, id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt, &t.CompletedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *SQLiteRepo) GetMonthlyStats(month string) ([]models.MonthlyStat, error) {
	query := `
		SELECT strftime('%Y-%m', completed_at) AS month, COUNT(*) 
		FROM tasks 
		WHERE completed = 1 AND completed_at IS NOT NULL`

	var args []any
	if month != "" {
		query += ` AND strftime('%Y-%m', completed_at) = ?`
		args = append(args, month)
	}
	query += ` GROUP BY month ORDER BY month DESC`

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []models.MonthlyStat
	for rows.Next() {
		var s models.MonthlyStat
		if err := rows.Scan(&s.Month, &s.CompletedCount); err != nil {
			return nil, err
		}
		stats = append(stats, s)
	}
	return stats, nil
}
func (r *SQLiteRepo) ListByDate(date string) ([]models.Task, error) {
	query := `
		SELECT id, title, completed, created_at, completed_at
		FROM tasks
		WHERE date(created_at, 'localtime') = ?
		   OR date(completed_at, 'localtime') = ?
		ORDER BY completed ASC, id DESC`

	rows, err := r.db.Query(query, date, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt, &t.CompletedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}
func (r *SQLiteRepo) GetByID(id int) (*models.Task, error) {
	var t models.Task
	err := r.db.QueryRow(
		`SELECT id, title, completed, created_at, completed_at FROM tasks WHERE id = ?`, id,
	).Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt, &t.CompletedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *SQLiteRepo) ListPrevious(today string) ([]models.Task, error) {
	query := `
		SELECT id, title, completed, created_at, completed_at
		FROM tasks
		WHERE date(created_at, 'localtime') < ?
		ORDER BY created_at DESC, id DESC`

	rows, err := r.db.Query(query, today)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt, &t.CompletedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}