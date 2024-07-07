package memdb

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/johnfercher/go-hexagonal/internal/core/consts/userstatus"
	"github.com/johnfercher/go-hexagonal/internal/core/models"
)

var ErrUserAlreadyCreated = errors.New("user already created")
var ErrUserNotFound = errors.New("user not found")

type UserRepository struct {
	db map[string]*models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: make(map[string]*models.User),
	}
}

func (m *UserRepository) CreatePending(ctx context.Context, creation *models.UserCreation) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	_, ok := m.db[id.String()]
	if ok {
		return "", ErrUserAlreadyCreated
	}

	m.db[id.String()] = &models.User{
		ID:        id.String(),
		CitizenID: creation.CitizenID,
		Name:      creation.Name,
		Address:   creation.Address,
		Status:    userstatus.Pending,
	}

	return id.String(), nil
}

func (m *UserRepository) UpdateStatus(ctx context.Context, id string, status userstatus.Status, info map[string]string) error {
	obj, ok := m.db[id]
	if !ok {
		return ErrUserNotFound
	}

	obj.Status = status
	m.db[id] = obj

	return nil
}
