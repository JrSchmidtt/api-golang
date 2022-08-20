package models

import (
	"postgressql/db"
)

func GetAll(id int) (todos []Todo, err error){
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM todos`, id)
	if err != nil {
		return
	}
	for rows.Next(){
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}
	return
}