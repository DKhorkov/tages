package ratelimiters

import (
	"context"
	"errors"
)

type Limiter struct {
	limit   int
	channel chan int
}

func (l *Limiter) Limit(ctx context.Context) error {
	if len(l.channel) >= l.limit {
		return errors.New("too many connections")
	}

	l.channel <- 1
	return nil
}

func NewLimiter(limit int, channel chan int) *Limiter {
	return &Limiter{limit: limit, channel: channel}
}
