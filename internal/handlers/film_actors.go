package handlers

import (
	"net/http"

	"github.com/AlexKorchagin/vk-films/internal/repo"
)

type FilmActorsHandlers struct {
	Repo *repo.Repository
}

func (a *FilmActorsHandlers) GetAllFilmActors() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
