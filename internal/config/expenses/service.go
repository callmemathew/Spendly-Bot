package expenses

import (
	"time"

	"Spendly-bot/internal/config/storage"
)

type Service struct {
	storage *storage.Storage
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) ParseText(text string) []ParsedExpense {
	return Parse(text)
}

func (s *Service) SaveParsedExpenses(userID int64, items []ParsedExpense) error {
	now := time.Now()
	date := now.Format("2006-01-02")

	for _, item := range items {
		expense := storage.Expense{
			UserID:      userID,
			Description: item.Description,
			Amount:      item.Amount,
			Category:    item.Category,
			Emoji:       item.Emoji,
			Currency:    "MDL",
			Date:        date,
		}

		if err := s.storage.SaveExpense(expense); err != nil {
			return err
		}
	}

	return nil
}
