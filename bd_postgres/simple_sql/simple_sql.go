package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	// по умолчанию используется схема public

	//создание таблицы
	sqlQuery := `
	CREATE TABLE tasks (
		id SERIAL PRIMARY KEY,
		title WARCHAR(200) NOT NULL,
		description WARCHAR(1000) NOT NULL,
		completed BOOLEAN NOT NULl,
		created_at TIMESTAMP NOT NULL,
		completet_at TIMESTAMP
	);
	`

	_, err := conn.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}

	return nil
}
