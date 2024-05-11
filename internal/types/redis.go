package types

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisDBT struct {
	DB  *redis.Client
	Ctx context.Context
}
