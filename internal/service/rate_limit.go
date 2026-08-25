package service

import (
	"context"
	"fmt"
)

func (s *NotesService) RateLimit(ctx context.Context, rpm int, ip string) (int, bool, error) {
	if rpm < 0 {
		return 0, false, fmt.Errorf("rate limit incorrect rpm")
	}
	if ip == "" {
		return 0, false, fmt.Errorf("rate limit: incorrect ip")
	}

	retryAfter, ok, err := s.cache.RateLimit(ctx, rpm, ip)
	if err != nil {
		return 0, false, fmt.Errorf("rate limit: %w", err)
	}

	return retryAfter, ok, nil
}
