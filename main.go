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

	if err := simple_sql.InsertRow(ctx, conn, simple_sql.TaskModel{
		Title:       "Учеба",
		Description: "Выучиться на разработчика",
		Completed:   false,
		Created_at:  time.Now(),
	}); err != nil {
		log.Println("error: ", err) //создание новой таблицы и получение из нее ошибки
	}

	if err := simple_sql.UpdateRow(ctx, conn); err != nil {
		log.Println("error: ", err) // смена статуса выполнения по id на true
	}

	if err := simple_sql.DeleteRow(ctx, conn, []int{1, 2, 3}); err != nil { //удаление по ИД
		log.Println("error: ", err)
	}

	tasks, err := simple_sql.SelectRows(ctx, conn) //все задачи
	if err != nil {
		log.Println("error: ", err)
	}

	fmt.Println(tasks)

	for _, task := range tasks {
		if task.ID == 3 {
			task.Title = "pay mobile"
			task.Description = "pay mobile operator"
			task.Completed = false

			time := time.Now()
			task.Completed_at = &time //от результата функции нельзя взять указатель

			if err := simple_sql.UpdateTask(ctx, conn, task); err != nil { //передаю изменения выбранной
				// таски в функции для смены параметров базы данных
				log.Println("error: ", err)
			}

			break
		}

	}

	fmt.Println("succeed")

}
