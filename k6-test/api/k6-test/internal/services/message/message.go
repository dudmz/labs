package message

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/carlosdamazio/k6-test/internal/models"
	"github.com/carlosdamazio/k6-test/internal/rest/entities"
)

type MessageService interface {
	List() []*models.Message
	Get(ID string) (*models.Message, error)
	Post(reqData *entities.PostMessageRequest) *models.Message
	Patch(ID, content string) (*models.Message, error)
	Delete(ID string) error
}

type Service struct {
	db map[string]*models.Message
}

func New(db map[string]*models.Message) *Service {
	return &Service{db: db}
}

func (s *Service) List() []*models.Message {
	messages := make([]*models.Message, 0)
	for _, msg := range s.db {
		messages = append(messages, msg)
	}
	return messages
}

func (s *Service) Get(ID string) (*models.Message, error) {
	if instance, ok := s.db[ID]; !ok {
		return nil, fmt.Errorf("get: instance %s not found", ID)
	} else {
		return instance, nil
	}
}

func (s *Service) Post(reqData *entities.PostMessageRequest) *models.Message {
	// avoid collision if UUID already exists
	id := uuid.NewString()
	for _, ok := s.db[id]; ok; {
		id = uuid.NewString()
	}

	instance := &models.Message{
		ID:      id,
		Content: reqData.Content,
		Length:  uint64(len(reqData.Content)),
	}
	s.db[id] = instance
	return instance
}

func (s *Service) Patch(ID, content string) (*models.Message, error) {
	if instance, ok := s.db[ID]; !ok {
		return nil, fmt.Errorf("patch: message of ID %s not found", ID)
	} else {
		instance.Content = content
		instance.Length = uint64(len(instance.Content))
		return instance, nil
	}
}

func (s *Service) Delete(ID string) error {
	if _, ok := s.db[ID]; !ok {
		return fmt.Errorf("delete: message of ID %s not found", ID)
	} else {
		delete(s.db, ID)
		return nil
	}
}
