package main

import (
	"bd/bd_postgres/simple_connection"
	"bd/bd_postgres/simple_sql"
	"context"
	"fmt"
	"log"
	"time"
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

	if err := simple_sql.InsertRow(ctx, conn, "Дворик3", "Поиграть во дворе", false, time.Now()); err != nil {
		log.Println("error: ", err) //создание новой таблицы и получение из нее ошибки
	}

	if err := simple_sql.UpdateRow(ctx, conn); err != nil {
		log.Println("error: ", err) // смена статуса выполнения по id на true
	}

	if err := simple_sql.DeleteRow(ctx, conn); err != nil {
		log.Println("error: ", err)
	}

	fmt.Println("succeed")
}
