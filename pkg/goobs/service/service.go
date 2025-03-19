package service

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/trace"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) SomeAction(ctx context.Context, sp trace.Span, name string) {
	sp.AddEvent("start some action")
	defer sp.AddEvent("finish some action")

	time.Sleep(10 * time.Millisecond)
}
