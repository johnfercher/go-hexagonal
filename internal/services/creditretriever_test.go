package services_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/johnfercher/go-hexagonal/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestNewCreditRetriever(t *testing.T) {
	t.Run("when construct, then build correctly", func(t *testing.T) {
		// Act
		sut := services.NewCreditRetriever()

		// Assert
		assert.NotNil(t, sut)
		assert.Equal(t, "*services.CreditRetriever", fmt.Sprintf("%T", sut))
	})
}

func TestCreditRetriever_Retrieve(t *testing.T) {
	t.Run("when retrieve, then retrieve correctly", func(t *testing.T) {
		// Arrange
		sut := services.NewCreditRetriever()

		// Act
		info := sut.Retrieve(context.TODO(), "citizen_id")

		// Assert
		assert.Len(t, info, 1)
		assert.Equal(t, "95%", info["score"])
	})
}
