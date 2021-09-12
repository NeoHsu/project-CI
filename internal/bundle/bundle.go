package bundle

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/NeoHsu/project-ci/internal/handlers"
	"github.com/NeoHsu/project-ci/internal/health"
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
				go func() {
					if err := h.Engine.Run(":9000"); err != nil {
						if errors.Is(err, http.ErrServerClosed) {
							fmt.Println("Shutting down the Application")
						} else {
							fmt.Printf("Error to Start Application: %v", err)
						}
					}
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping application")
				return nil
			},
		},
	)
}
