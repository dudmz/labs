package message

import "github.com/carlosdamazio/k6-test/internal/models"

type MessageService interface {
	List() []*models.Message
	Get(ID string) *models.Message
	Post() error
	Patch() (*models.Message, error)
	Delete(ID string) error
}

type Service struct {
	db map[string]*models.Message
}

func New(db map[string]*models.Message) *Service {
	return &Service{db: db}
}

func (s *Service) List() []*models.Message {
	messages := make([]*models.Message, len(s.db))
	for _, msg := range s.db {
		messages = append(messages, msg)
	}
	return messages
}

func (s *Service) Get(ID string) *models.Message {
	if _, ok := s.db[ID]; !ok {
		return nil
	}

	return s.db[ID]
}

func (s *Service) Post() error {
	return nil
}

func (s *Service) Patch() (*models.Message, error) {
	return nil, nil
}

func (s *Service) Delete(ID string) error {
	return nil
}
