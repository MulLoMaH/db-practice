package simple_sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, conn *pgx.Conn) ([]TaskModel, error) {
	sqlQuery := `
	SELECT id, title, description, completed, created_at, completed_at
	FROM tasks
	WHERE completed = FALSE;
	`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []TaskModel

	for rows.Next() {
		var task TaskModel

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.Created_at,
			&task.Completed_at,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)

		printTask(task)

	}

	return tasks, nil
}

func printTask(task TaskModel) { //красивый вывод таск из БД
	fmt.Println("-----------------------------")
	fmt.Println("id: ", task.ID)
	fmt.Println("title: ", task.Title)
	fmt.Println("description: ", task.Description)
	fmt.Println("completed: ", task.Completed)
	fmt.Println("created ad: ", task.Created_at)
	fmt.Println("completed at: ", task.Completed_at)
}

// func SelectRows(ctx context.Context, conn *pgx.Conn) error {
// 	sqlQuery := `
// 	SELECT id, title, description, completed, created_at, completet_at
// 	FROM tasks
// 	WHERE completed = FALSE;
// 	`

// 	// sqlQuery := ` - создаем запрос
// 	// SELECT id, title, description, completed, created_at, completed_at - просим вернуть из
// 	// базы данных параметры стольбцов id, title, description, completed, created_at, completed_at
// 	// FROM tasks - из таблицы tasks
// 	// WHERE completed = FALSE; // все строки у которых completed = FALSE
// 	// `

// 	rows, err := conn.Query(ctx, sqlQuery) //вернет все значения соответствующие условию
// 	if err != nil {
// 		return err
// 	}

// 	defer rows.Close() // отложенное освобождение подключения к базе данных

// 	for rows.Next() { //в цикле читаем значения полученные из conn.Query до
// 		// конца (.Next() вернет по окончанию получения значений false и цикл остановиться)

// 		// создаем переменные в которые получим значения
// 		var id int
// 		var title string
// 		var description string
// 		var completed bool
// 		var created_at time.Time
// 		var completed_at *time.Time

// 		err := rows.Scan( //сканируем в каждой
// 			// итерации полученные значения в наши созданные переменные по указателю
// 			&id,
// 			&title,
// 			&description,
// 			&completed,
// 			&created_at,
// 			&completed_at, //указатель на указатель
// 		)

// 		if err != nil {
// 			return err
// 		}

// 		fmt.Println(id, title, description, completed, created_at, completed_at) //выводим на экран результат
// 	}

// 	//conn.QueryRow() вернет вервую подходящую строку под условие

// 	return nil
// }
