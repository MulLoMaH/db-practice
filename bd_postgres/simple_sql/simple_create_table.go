package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	UPDATE tasks
	SET description = ':0'
	WHERE completed = false;
	`

	_, err := conn.Exec(ctx, sqlQuery)

	return err
}

func UpdateTask(
	ctx context.Context,
	conn *pgx.Conn,
	task TaskModel,
) error {
	sqlQuery := `
	UPDATE tasks
	SET title=$1, decription=$2, completed=$3, created_at=$4, copleted_at=$5
	WHERE id=$6 
	`

	_, err := conn.Exec(
		ctx,
		sqlQuery,
		task.Title,
		task.Description,
		task.Completed,
		task.Created_at,
		task.Completed_at,
		task.ID,
	)

	return err
}

// sqlQuery := ` - создаем логику - условие
// UPDATE tasks - изменить таблицу таск
// SET title=$1, decription=$2, completed=$3, created_at=$4, copleted_at=$5 - изменить набор
// полей <- (всех полей кроме ИД)
// WHERE id=%6 - у строки с ИД передаваемое значение
// `

// sqlQuery := ` //создаем действие для таблицы
// UPDATE tasks // указываем что будут изменения таблицы tasks
// SET completed = TRUE //указываем что будут изменения поля completed на true
// WHERE id = 3 // условие id = 3
// `
