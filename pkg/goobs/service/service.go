package service

import (
	"context"
	"goobs/tracer"
	"time"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	tracer trace.Tracer
	trSvc  trace.Tracer
}

func New(tr trace.Tracer) *Service {
	traceSvc, err := tracer.InitTracer(context.Background(), "goobs-service")
	if err != nil {
		panic(err)
	}
	trSvc := traceSvc.Tracer("goobs-tracer-service")

	return &Service{
		tracer: tr,
		trSvc:  trSvc,
	}
}

func (s *Service) SomeAction(ctx context.Context, name string, data map[string]string) {
	_, sp := s.tracer.Start(ctx, "some-action")
	defer sp.End()

	sp.AddEvent("start some action")
	defer sp.AddEvent("finish some action")

	p := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
	ctx = p.Extract(ctx, propagation.MapCarrier(data))

	_, span := s.trSvc.Start(ctx, "ping-receive")
	defer span.End()

	time.Sleep(10 * time.Millisecond)
}
