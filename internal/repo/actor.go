package repo

import (
	"github.com/AlexKorchagin/vk-films/gen"
	"github.com/jmoiron/sqlx"
)

type actorRepo struct {
	conn *sqlx.DB
}

func newActorRepo(conn *sqlx.DB) *actorRepo {
	return &actorRepo{
		conn: conn,
	}
}

func SaveActorToDatabase(a gen.Actor, r *Repository) error {
	_, err := r.conn.Exec("INSERT INTO actors (DateOfBirth, Gender, LastName, Name) VALUES ($1, $2, $3, $4)",
		a.DateOfBirth, a.Gender, a.LastName, a.Name)
	if err != nil {
		return err
	}

	return nil
}

func UpdateActorDB(s string, r *Repository) error {
	_, err := r.conn.Exec(s)
	if err != nil {
		return err
	}
	return nil
}

func DeleteActorByID(id int64, r *Repository) error {
	_, err := r.conn.Exec("DELETE FROM actors WHERE ID = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func GetAllActors(r *Repository) ([]gen.Actor, error) {
	rows, err := r.conn.Query("SELECT * FROM actors")
	if err != nil {
		return nil, err
	}
	var actors []gen.Actor

	for rows.Next() {
		var actor gen.Actor
		err := rows.Scan(&actor.ID, &actor.Name, &actor.LastName, &actor.DateOfBirth, &actor.Gender)
		if err != nil {
			return nil, err
		}
		// Добавление фильма в слайс
		actors = append(actors, actor)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return actors, nil
}
