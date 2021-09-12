package bundle

import (
	"context"
	"fmt"

	"github.com/NeoHsu/project-CI/internal/handlers"
	"github.com/NeoHsu/project-CI/internal/health"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handlers.Module,
	health.Module,
	fx.Invoke(registerHooks),
)

func registerHooks(lifecycle fx.Lifecycle, h *handlers.HttpHandler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				fmt.Println("Starting application on 9000")
				go h.Engine.Run(":9000")
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping application")
				return nil
			},
		},
	)
}
