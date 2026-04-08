package simple_sql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func InsertRow(
	ctx context.Context,
	conn *pgx.Conn,
	title string,
	description string,
	completed bool,
	createdAt time.Time,
) error { //добавляем строки в БД
	sqlQuery := `
	INSERT INTO tasks (title, description, completed, created_at)
	VALUES ($1, $2, $3, $4);
	` //указываем порядковый номер передаваемых параметров через $ в conn.Exec(<-)

	//метод отправляет данные в БД по зарегистрированному маршруту
	_, err := conn.Exec(ctx, sqlQuery, title, description, completed, createdAt) //отправляем данные в таблицу //

	return err //если будет ошибка то вернется она если нил то нил
	//обработка в другой функции все равно
}

//  sqlQuery := `
// 	INSERT INTO tasks(title, descrption, completed, ompleted_at)(название таблицы можно со схемой или без в случае стандартного пути)
//  VALUES ('Домашка', 'Сделать домашку по матеше до 20.03.26', FALSE, '2025-11-26 18:00:05'); //таким образом заполняются поля в postgres
// id не заполняем так как он автоматический
// 	`
