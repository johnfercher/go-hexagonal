package kycservices

import (
	"context"

	"github.com/johnfercher/go-hexagonal/internal/core/consts/userstatus"
	"github.com/johnfercher/go-hexagonal/internal/core/models"
	"github.com/johnfercher/go-hexagonal/internal/core/ports"
)

type UserRegister struct {
	infoRetriever []ports.UserInfoRetriever
	repository    ports.UserRepository
}

func NewUserRegister(repository ports.UserRepository, retrievers ...ports.UserInfoRetriever) *UserRegister {
	return &UserRegister{
		repository:    repository,
		infoRetriever: retrievers,
	}
}

func (u *UserRegister) Register(ctx context.Context, creation *models.UserCreation) (string, error) {
	id, err := u.repository.CreatePending(ctx, creation)
	if err != nil {
		return "", err
	}

	info := make(map[string]string)
	for _, retriever := range u.infoRetriever {
		iInfo := retriever.Retrieve(ctx, creation.CitizenID)
		info = u.merge(info, iInfo)
	}

	status := userstatus.Allowed
	if len(info) > 0 {
		status = userstatus.Restricted
	}
	if len(info) > 1 {
		status = userstatus.Denied
	}

	return id, u.repository.UpdateStatus(ctx, id, status, info)
}

func (u *UserRegister) merge(a map[string]string, b map[string]string) map[string]string {
	for key, value := range b {
		a[key] = value
	}

	return a
}
