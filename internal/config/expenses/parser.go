package expenses

import (
	"strconv"
	"strings"
)

func Parse(text string) []ParsedExpense {
	words := strings.Fields(text)

	var result []ParsedExpense
	var descriptionWords []string

	for _, word := range words {
		cleanWord := cleanAmountWord(word)

		amount, err := strconv.ParseFloat(cleanWord, 64)
		if err == nil {
			if len(descriptionWords) == 0 {
				continue
			}

			description := strings.Join(descriptionWords, " ")
			category, emoji := detectCategory(description)

			result = append(result, ParsedExpense{
				Description: normalizeDescription(description),
				Amount:      amount,
				Category:    category,
				Emoji:       emoji,
			})

			descriptionWords = []string{}
			continue
		}

		descriptionWords = append(descriptionWords, word)
	}

	return result
}

func cleanAmountWord(word string) string {
	word = strings.TrimSpace(word)
	word = strings.Trim(word, ".,!?;:")
	word = strings.ReplaceAll(word, ",", ".")

	return word
}

func detectCategory(description string) (string, string) {
	text := strings.ToLower(description)

	switch {
	case containsAny(text, []string{"такси", "метро", "маршрутка", "автобус", "бензин", "топливо"}):
		return "Transport", "🚕"

	case containsAny(text, []string{"овощи", "фрукты", "вода", "хлеб", "молоко", "мясо", "продукты", "еда"}):
		return "Food", "🥦"

	case containsAny(text, []string{"врач", "окулист", "аптека", "таблетки", "анализы", "стоматолог"}):
		return "Health", "🏥"

	case containsAny(text, []string{"эндис", "кафе", "кофе", "шавуха", "ресторан", "пицца", "бургер"}):
		return "Cafe", "🍔"

	case containsAny(text, []string{"подписка", "spotify", "netflix", "icloud", "youtube", "apple"}):
		return "Subscriptions", "💳"

	case containsAny(text, []string{"зал", "спорт", "теннис", "протеин", "креатин"}):
		return "Sport", "🏋️"

	case containsAny(text, []string{"ремонт", "квартира", "дом", "краны", "дверь"}):
		return "Home", "🏠"

	default:
		return "Other", "💸"
	}
}

func containsAny(text string, words []string) bool {
	for _, word := range words {
		if strings.Contains(text, word) {
			return true
		}
	}

	return false
}

func normalizeDescription(description string) string {
	description = strings.TrimSpace(description)

	if description == "" {
		return description
	}

	runes := []rune(description)
	runes[0] = []rune(strings.ToUpper(string(runes[0])))[0]

	return string(runes)
}
