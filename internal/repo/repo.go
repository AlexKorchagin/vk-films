package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type Repository struct {
	conn  *sqlx.DB
	Actor *actorRepo
	Film  *filmRepo
}

func NewRepository() (*Repository, error) {

	db, err := NewPostgresDB(Config{
		Host:     viper.GetString("Host"),
		Port:     viper.GetString("DBPort"),
		Username: viper.GetString("Username"),
		Password: viper.GetString("Password"),
		DBName:   viper.GetString("DBName"),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to initialize db: %s", err.Error())
	}

	actorRepo := newActorRepo(db)
	filmRepo := newFilmRepo(db)

	return &Repository{
		conn:  db,
		Actor: actorRepo,
		Film:  filmRepo,
	}, nil
}
