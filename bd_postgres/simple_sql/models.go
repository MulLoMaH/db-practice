package simple_sql

import "time"

type TaskModel struct { //модель создается для общения между базой
	// данных и Го приложением как DTO в http
	ID           int
	Title        string
	Description  string
	Completed    bool
	Created_at   time.Time
	Completed_at *time.Time
}
