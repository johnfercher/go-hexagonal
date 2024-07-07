package ports

import (
	"context"
	"github.com/johnfercher/go-hexagonal/internal/core/models"
)

type UserRegister interface {
	Register(ctx context.Context, creation *models.UserCreation) (string, error)
}
