package adapters

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	rdb *redis.Client
}

func NewRedis(host string) *Redis {
	addr := host + ":6379"
	return &Redis{rdb: redis.NewClient(&redis.Options{Addr: addr})}
}

func (r *Redis) RateLimit(ctx context.Context, rpm int, ip string) (int, bool, error) {
	key := fmt.Sprintf("limit:%s", ip)

	res, err := r.rdb.Incr(ctx, key).Result()
	if err != nil {
		return 0, false, fmt.Errorf("rate limit: %w", err)
	}
	if res == 1 {
		r.rdb.Expire(ctx, key, 60*time.Second)
	}

	if res > int64(rpm) {
		ttl, err := r.rdb.TTL(ctx, key).Result()
		if err != nil {
			return 0, false, fmt.Errorf("getting ttl: %w", err)
		}
		return int(ttl.Seconds()), false, nil
	}

	return 0, true, nil
}
