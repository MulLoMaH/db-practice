package main

import (
	"bd/bd_postgres/simple_connection"
	"bd/bd_postgres/simple_sql"
	"context"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()

	conn, err := simple_connection.CreateConnection(ctx) //вызов функции подключения к базе данных
	if err != nil {
		log.Println("error: ", err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil { //передаем подключение в функцию
		log.Println("error: ", err)
	}

	fmt.Println("Таблицав базе данных была создана успешно")
}
