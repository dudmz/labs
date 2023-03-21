package db

import "github.com/carlosdamazio/k6-test/internal/models"

var (
	Messages map[string]*models.Message = make(map[string]*models.Message)
)
