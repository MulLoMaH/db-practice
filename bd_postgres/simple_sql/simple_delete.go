package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	DELETE FROM tasks
	WHERE id = 5;
	`
	_, err := conn.Exec(ctx, sqlQuery)

	return err

}

// sqlQuery := ` - sql запрос
// 	DELETE FROM tasks - удалить стороки таск (без условия ниже - удаляются все )
// 	WHERE id = 8; - условие если id равен 8
// 	`
