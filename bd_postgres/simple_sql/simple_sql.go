package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	// по умолчанию используется схема public

	//создание таблицы
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(200) NOT NULL,
		description VARCHAR(1000) NOT NULL,
		completed BOOLEAN NOT NULl,
		created_at TIMESTAMP NOT NULL,
		completet_at TIMESTAMP,


		UNIQUE(title)
	);
	`

	//метод отправляет данные в БД по зарегистрированному маршруту
	_, err := conn.Exec(ctx, sqlQuery) //вызов функции принимающей контекст и таблицу
	//  и переменное число разных аргументов создающую таблицу по адресу подключения
	if err != nil {
		return err
	}

	return nil
}

// sqlQuery := `
// 	CREATE TABLE IF NOT EXISTS tasks ( - создать таблицу если она еще не создана
// 		id SERIAL PRIMARY KEY, - поле ИД автоматически заполняемое поле ИД ключ начинается с 1 и ++
// 		title VARCHAR(200) NOT NULL, - заголовок СТРОКА(до 200 символов) не может быть пустым
// 		description VARCHAR(1000) NOT NULL, - тело задачи СТРОКА(до 100 символов) не может быть пустым
// 		completed BOOLEAN NOT NULl, - выполнена или нет БУЛЕАН типа не может быть пустым
// 		created_at TIMESTAMP NOT NULL, - время создания задачи типа ВРЕМЯ не может быть пустым
// 		completet_at TIMESTAMP - время выполнения задачи типа ВРЕМЯ может быть пустым
//
//
//    UNIQUE(title) - обозначает уникальность заданного столбика (в нашем случае выбран заголовок)
// 	); - ; обязательна и после последнего заголовка таблицы не ставится запятая
// 	`
