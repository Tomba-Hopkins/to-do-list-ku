package models

import (
	"to-do-list/config"
	"to-do-list/entities"
)

func GetAll() []entities.Todo {
	rows, err := config.DB.Query(`SELECT * FROM todolist`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var todos []entities.Todo

	for rows.Next() {
		var todo entities.Todo

		if err := rows.Scan(&todo.Id, &todo.Todoco, &todo.Todoce, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			panic(err)
		}

		todos = append(todos, todo)
	}

	return todos
}

func CreateTodo(todo entities.Todo) bool {
	result, err := config.DB.Exec(`INSERT INTO todolist (todoco, todoce, created_at, updated_at) VALUE (?, ?, ?, ?)`, todo.Todoco, todo.Todoce, todo.CreatedAt, todo.UpdatedAt)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM todolist WHERE id = ?`, id)
	return err
}