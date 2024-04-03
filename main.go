package main

import (
	"awesomeProject/internal/repositories"
	"awesomeProject/providers"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			zap.NewExample,
			providers.NewServer,
			repositories.NewTaskRepository,
		),
		fx.Invoke(providers.NewRouter),
	).Run()
}
