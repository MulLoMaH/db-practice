package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn) error { //добавляем строки в БД
	sqlQuery := `
	INSERT INTO tasks (title, description, completed, created_at)
	VALUES ('Домашка', 'Сделать домашку по матеше до 20.03.26', FALSE, '2025-11-26 18:00:05');
	`

	//метод отправляет данные в БД по зарегистрированному маршруту
	_, err := conn.Exec(ctx, sqlQuery) //отправляем данные в таблицус

	return err //если будет ошибка то вернется она если нил то нил
	//обработка в другой функции все равно
}

//  sqlQuery := `
// 	INSERT INTO tasks(title, descrption, completed, ompleted_at)(название таблицы можно со схемой или без в случае стандартного пути)
//  VALUES ('Домашка', 'Сделать домашку по матеше до 20.03.26', FALSE, '2025-11-26 18:00:05'); //таким образом заполняются поля в postgres
// id не заполняем так как он автоматический
// 	`
