package models

import (
	"to-do-list/config"
	"to-do-list/entities"
)

func GetAllResol() []entities.Resol {
	rows, err := config.DB.Query(`SELECT * FROM resol`)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var resols []entities.Resol

	for rows.Next() {
		var resol entities.Resol

		if err := rows.Scan(&resol.IdR, &resol.Resolco, &resol.Resolce, &resol.CreatedAt, &resol.UpdatedAt); err != nil {
			panic(err)
		}

		resols = append(resols, resol)
	}

	return resols
}


func CreateResol(resol entities.Resol) bool {
	result, err := config.DB.Exec(`INSERT INTO resol (resolco, resolce, created_at, updated_at) VALUE (?, ?, ?, ?)`, resol.Resolco, resol.Resolce, resol.CreatedAt, resol.UpdatedAt)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func DeleteResol(id int) error {
	_, err := config.DB.Exec(`DELETE FROM resol WHERE id = ?`, id)
	return err
}