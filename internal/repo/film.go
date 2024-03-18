package repo

import (
	"fmt"

	"github.com/AlexKorchagin/vk-films/gen"
	"github.com/jmoiron/sqlx"
)

type filmRepo struct {
	conn *sqlx.DB
}

func newFilmRepo(conn *sqlx.DB) *filmRepo {
	return &filmRepo{
		conn: conn,
	}
}

func SaveFilmToDatabase(f gen.Film, r *Repository) error {
	_, err := r.conn.Exec("INSERT INTO films (PublishDay, Description, Rating, Name) VALUES ($1, $2, $3, $4)",
		f.PublishDay, f.Description, f.Rating, f.Name)
	if err != nil {
		return err
	}

	return nil
}

func UpdateFilmDB(s string, r *Repository) error {
	_, err := r.conn.Exec(s)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFilmByID(id int64, r *Repository) error {
	_, err := r.conn.Exec("DELETE FROM films WHERE ID = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func GetAllFilms(sortBy string, r *Repository) ([]gen.Film, error) {
	query := fmt.Sprintf("SELECT * FROM films ORDER BY %s", sortBy)

	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, err
	}

	var films []gen.Film

	for rows.Next() {
		var film gen.Film
		err := rows.Scan(&film.ID, &film.Name, &film.Description, &film.PublishDay, &film.Rating)
		if err != nil {
			return nil, err
		}
		// Добавление фильма в слайс
		films = append(films, film)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return films, nil
}
