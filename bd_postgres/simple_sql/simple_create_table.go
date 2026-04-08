package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	UPDATE tasks
	SET completed = FALSE
	WHERE id = 1;
	`

	_, err := conn.Exec(ctx, sqlQuery)

	return err
}

// sqlQuery := ` //создаем действие для таблицы
// UPDATE tasks // указываем что будут изменения таблицы tasks
// SET completed = TRUE //указываем что будут изменения поля completed на true
// WHERE id = 3 // условие id = 3
// `
