package admin

type Service struct{}

func NewService() *Service { return &Service{} }

func (s *Service) FindAll() []string { return []string{"example"} }
