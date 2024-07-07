package ports

import (
	"context"
)

type UserInfoRetriever interface {
	Retrieve(ctx context.Context, citizenID string) map[string]string
}
