package main

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.TimeKey = ""

	cfg.Sampling = &zap.SamplingConfig{
		Initial:    3,
		Thereafter: 3,
		Hook: func(e zapcore.Entry, sd zapcore.SamplingDecision) {
			if sd == zapcore.LogDropped {
				fmt.Println("event dropped...")
			}
		},
	}

	logger, _ := cfg.Build()
	zap.ReplaceGlobals(logger)
}

func main() {
	for i := 1; i <= 10; i++ {
		zap.S().Infow("Testing sampling", "index", i)
	}
}
