package kycservices_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/johnfercher/go-hexagonal/internal/core/consts/userstatus"
	"github.com/johnfercher/go-hexagonal/internal/fixture"
	"github.com/johnfercher/go-hexagonal/internal/kycservices"
	"github.com/johnfercher/go-hexagonal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewUserRegister(t *testing.T) {
	t.Run("when create, then create correctly", func(t *testing.T) {
		// Act
		sut := kycservices.NewUserRegister(nil)

		// Assert
		assert.NotNil(t, sut)
		assert.Equal(t, "*kycservices.UserRegister", fmt.Sprintf("%T", sut))
	})
}

// nolint:dupl // testing each case
func TestUserRegister_Register(t *testing.T) {
	t.Run("when cannot create pending, then return error", func(t *testing.T) {
		// Arrange
		ctx := context.TODO()
		creation := fixture.UserCreation()
		errToReturn := fixture.Error()

		repository := mocks.NewUserRepository(t)
		repository.EXPECT().CreatePending(ctx, creation).Return("", errToReturn)

		sut := kycservices.NewUserRegister(repository)

		// Act
		id, err := sut.Register(ctx, creation)

		// Assert
		assert.Empty(t, id)
		assert.Equal(t, errToReturn, err)
	})
	t.Run("when create pending but cannot update, then return error", func(t *testing.T) {
		// Arrange
		ctx := context.TODO()
		creation := fixture.UserCreation()
		uid, _ := uuid.NewRandom()
		info1 := fixture.Info(0)
		info2 := fixture.Info(0)
		infoMerged := merge(info1, info2)
		errToReturn := fixture.Error()

		repository := mocks.NewUserRepository(t)
		repository.EXPECT().CreatePending(ctx, creation).Return(uid.String(), nil)
		repository.EXPECT().UpdateStatus(ctx, uid.String(), userstatus.Allowed, infoMerged).Return(errToReturn)

		retriever1 := mocks.NewUserInfoRetriever(t)
		retriever1.EXPECT().Retrieve(ctx, creation.CitizenID).Return(info1)

		retriever2 := mocks.NewUserInfoRetriever(t)
		retriever2.EXPECT().Retrieve(ctx, creation.CitizenID).Return(info2)

		sut := kycservices.NewUserRegister(repository, retriever1, retriever2)

		// Act
		id, err := sut.Register(ctx, creation)

		// Assert
		assert.Equal(t, uid.String(), id)
		assert.Equal(t, errToReturn, err)
	})
	t.Run("when has one info, then create as allowed", func(t *testing.T) {
		// Arrange
		ctx := context.TODO()
		creation := fixture.UserCreation()
		uid, _ := uuid.NewRandom()
		info1 := fixture.Info(0)
		info2 := fixture.Info(0)
		infoMerged := merge(info1, info2)

		repository := mocks.NewUserRepository(t)
		repository.EXPECT().CreatePending(ctx, creation).Return(uid.String(), nil)
		repository.EXPECT().UpdateStatus(ctx, uid.String(), userstatus.Allowed, infoMerged).Return(nil)

		retriever1 := mocks.NewUserInfoRetriever(t)
		retriever1.EXPECT().Retrieve(ctx, creation.CitizenID).Return(info1)

		retriever2 := mocks.NewUserInfoRetriever(t)
		retriever2.EXPECT().Retrieve(ctx, creation.CitizenID).Return(info2)

		sut := kycservices.NewUserRegister(repository, retriever1, retriever2)

		// Act
		id, err := sut.Register(ctx, creation)

		// Assert
		assert.Equal(t, uid.String(), id)
		assert.Nil(t, err)
	})
	t.Run("when has two info, then create as restricted", func(t *testing.T) {
		// Arrange
		ctx := context.TODO()
		creation := fixture.UserCreation()
		uid, _ := uuid.NewRandom()
		info1 := fixture.Info(1)
		info2 := fixture.Info(0)
		infoMerged := merge(info1, info2)

		repository := mocks.NewUserRepository(t)
		repository.EXPECT().CreatePending(ctx, creation).Return(uid.String(), nil)
		repository.EXPECT().UpdateStatus(ctx, uid.String(), userstatus.Restricted, infoMerged).Return(nil)

		retriever1 := mocks.NewUserInfoRetriever(t)
		retriever1.EXPECT().Retrieve(ctx, creation.CitizenID).Return(info1)

		retriever2 := mocks.NewUserInfoRetriever(t)
		retriever2.EXPECT().Retrieve(ctx, creation.CitizenID).Return(info2)

		sut := kycservices.NewUserRegister(repository, retriever1, retriever2)

		// Act
		id, err := sut.Register(ctx, creation)

		// Assert
		assert.Equal(t, uid.String(), id)
		assert.Nil(t, err)
	})
	t.Run("when has two info, then create as restricted", func(t *testing.T) {
		// Arrange
		ctx := context.TODO()
		creation := fixture.UserCreation()
		uid, _ := uuid.NewRandom()
		info1 := fixture.Info(1)
		info2 := fixture.Info(2)
		infoMerged := merge(info1, info2)

		repository := mocks.NewUserRepository(t)
		repository.EXPECT().CreatePending(ctx, creation).Return(uid.String(), nil)
		repository.EXPECT().UpdateStatus(ctx, uid.String(), userstatus.Denied, infoMerged).Return(nil)

		retriever1 := mocks.NewUserInfoRetriever(t)
		retriever1.EXPECT().Retrieve(ctx, creation.CitizenID).Return(info1)

		retriever2 := mocks.NewUserInfoRetriever(t)
		retriever2.EXPECT().Retrieve(ctx, creation.CitizenID).Return(info2)

		sut := kycservices.NewUserRegister(repository, retriever1, retriever2)

		// Act
		id, err := sut.Register(ctx, creation)

		// Assert
		assert.Equal(t, uid.String(), id)
		assert.Nil(t, err)
	})
}

func merge(a map[string]string, b map[string]string) map[string]string {
	for key, value := range b {
		a[key] = value
	}

	return a
}
