package main

import (
	"context"
	"courses/golang/inventory-project/database"
	"courses/golang/inventory-project/internal/repository"
	"courses/golang/inventory-project/internal/service"
	"courses/golang/inventory-project/settings"
	"fmt"

	"github.com/jmoiron/sqlx"
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

			func(s *settings.Settings) {
				fmt.Println(s)
			},

			func(db *sqlx.DB) {
				result, err := db.Query("SELECT * FROM USERS")
				if err != nil {
					panic(err)
				}
				fmt.Println(result.Columns())
			},
		),
	)

	app.Run()
}
