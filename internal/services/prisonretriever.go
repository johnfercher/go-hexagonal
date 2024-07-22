package services

import (
	"context"
)

type PrisonRetriever struct{}

func NewPrisonRetriever() *PrisonRetriever {
	return &PrisonRetriever{}
}

func (op *PrisonRetriever) Retrieve(ctx context.Context, citizenID string) map[string]string {
	m := make(map[string]string)

	m["sao_paulo"] = "3 days"

	return m
}
