package main

import (
	"context"
	"courses/golang/inventory-project/database"
	"courses/golang/inventory-project/internal/repository"
	"courses/golang/inventory-project/internal/service"
	"courses/golang/inventory-project/settings"

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
		),

		fx.Invoke(
			func(ctx context.Context) {
				s, _ := settings.New()

				db, err := database.New(ctx, s)

				if err != nil {
					println("ddd")
				}

				r := repository.New(db)

				sr := service.New(r)

				println(sr.Test())
			},
		),
	)

	app.Run()
}
