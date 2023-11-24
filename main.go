package main

import (
	"context"
	"courses/golang/inventory-project/database"
	"courses/golang/inventory-project/internal/api"
	"courses/golang/inventory-project/internal/repository"
	"courses/golang/inventory-project/internal/service"
	"courses/golang/inventory-project/settings"
	"fmt"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
			api.New,
			echo.New,
		),

		fx.Invoke(
			setLifeCycle,
		),
	)

	app.Run()
}

// The function sets up the lifecycle hooks for starting and stopping an API server.
func setLifeCycle(lc fx.Lifecycle, a *api.API, s *settings.Settings, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", s.Port)

			go a.Start(e, address)

			return nil
		},

		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
