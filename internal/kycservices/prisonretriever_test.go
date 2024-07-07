package kycservices_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/johnfercher/go-hexagonal/internal/kycservices"
	"github.com/stretchr/testify/assert"
)

func TestNewPrisonRetriever(t *testing.T) {
	t.Run("when construct, then build correctly", func(t *testing.T) {
		// Act
		sut := kycservices.NewPrisonRetriever()

		// Assert
		assert.NotNil(t, sut)
		assert.Equal(t, "*kycservices.PrisonRetriever", fmt.Sprintf("%T", sut))
	})
}

func TestPrisonRetriever_Retrieve(t *testing.T) {
	t.Run("when retrieve, then retrieve correctly", func(t *testing.T) {
		// Arrange
		sut := kycservices.NewPrisonRetriever()

		// Act
		info := sut.Retrieve(context.TODO(), "citizen_id")

		// Assert
		assert.Len(t, info, 1)
		assert.Equal(t, "3 days", info["sao_paulo"])
	})
}
