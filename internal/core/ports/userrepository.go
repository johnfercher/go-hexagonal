package ports

import (
	"context"
	"github.com/johnfercher/go-hexagonal/internal/core/consts/userstatus"
	"github.com/johnfercher/go-hexagonal/internal/core/models"
)

type UserRepository interface {
	CreatePending(ctx context.Context, creation *models.UserCreation) (string, error)
	UpdateStatus(ctx context.Context, id string, status userstatus.Status, m map[string]string) error
}
