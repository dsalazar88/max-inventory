package database

import (
	"context"
	"courses/golang/inventory-project/settings"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// The New function creates a new database connection using the provided settings.
func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		s.DB.User,
		s.DB.Password,
		s.DB.Host,
		s.DB.Port,
		s.DB.Name,
	)
	fmt.Println(connString)
	return sqlx.ConnectContext(ctx, "mysql", connString)
}
