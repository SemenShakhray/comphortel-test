package service

import (
	"comphortel-test/internal/delivery/handlers"
	"comphortel-test/internal/models"
	"comphortel-test/utils/logger/sl"
	"fmt"
	"log/slog"
	"strconv"
)

type Service struct {
	repo Repositorer
	log  *slog.Logger
}

func NewService(repo Repositorer, log *slog.Logger) handlers.Servicer {
	return &Service{
		repo: repo,
		log:  log,
	}
}

type Repositorer interface {
	GetUser(id int) (*models.User, error)
	Close() error
}

func (s *Service) GetUser(id string) (*models.User, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		s.log.Error("failed to convert id to int", sl.Err(err))

		return nil, fmt.Errorf("failed to convert id to int: %w", err)
	}

	return s.repo.GetUser(idInt)
}
