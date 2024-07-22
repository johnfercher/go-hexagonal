package services

import (
	"context"
)

type CreditRetriever struct{}

func NewCreditRetriever() *CreditRetriever {
	return &CreditRetriever{}
}

func (op *CreditRetriever) Retrieve(ctx context.Context, citizenID string) map[string]string {
	m := make(map[string]string)

	m["score"] = "95%"

	return m
}
