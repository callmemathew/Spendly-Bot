package expenses

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ParseText(text string) []ParsedExpense {
	return Parse(text)
}
