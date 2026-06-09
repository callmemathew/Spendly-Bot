package storage

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	DB *sql.DB
}
type Expense struct {
	ID          int64
	UserID      int64
	Description string
	Amount      float64
	Category    string
	Emoji       string
	Currency    string
	Date        string
	CreatedAt   string
}

func NewSQLiteStorage(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping sqlite database: %w", err)
	}

	storage := &Storage{
		DB: db,
	}

	if err := storage.createTables(); err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *Storage) createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		description TEXT NOT NULL,
		amount REAL NOT NULL,
		category TEXT NOT NULL,
		emoji TEXT NOT NULL,
		currency TEXT NOT NULL DEFAULT 'MDL',
		date TEXT NOT NULL,
		created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := s.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create expenses table: %w", err)
	}

	return nil
}

func (s *Storage) Close() error {
	return s.DB.Close()
}
func (s *Storage) SaveExpense(expense Expense) error {
	query := `
	INSERT INTO expenses (
		user_id,
		description,
		amount,
		category,
		emoji,
		currency,
		date
	) VALUES (?, ?, ?, ?, ?, ?, ?);
	`

	_, err := s.DB.Exec(
		query,
		expense.UserID,
		expense.Description,
		expense.Amount,
		expense.Category,
		expense.Emoji,
		expense.Currency,
		expense.Date,
	)

	if err != nil {
		return fmt.Errorf("failed to save expense: %w", err)
	}

	return nil
}
