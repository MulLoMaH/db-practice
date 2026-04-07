package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5" //библиотека для работы с базой данных
)

//"postgres//:YourUserName:YourPassword@YourHostName:5432/YourDatabaseName"

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {

	//(технология согдания подключения к базе данных PostgeSQL)
	//"postgres//:YourUserName(по умолчанию - postgres)
	// :YourPassword@YourHostName(без сервера - localhost):5432(порт на котором находится база данных)
	// /YourDatabaseName" (название подраздела - по умолчанию - postgres)
	// вернет указатель на подключение к базе данных и ошибку

	return pgx.Connect(ctx, "postgres://postgres:2906@localhost:5432/postgres") //подключение к базе данных
	// 	if err != nil {
	// 		log.Println("error: ", err)
	// 	}

	// 	if err := conn.Ping(ctx); err != nil { // тестовый запрос в БД для проверки доступа
	// 		log.Println("error: ", err)
	// 	}

	// 	fmt.Println("Подключение к базе данных произошло успешно!")

	// return conn
}
