package handlers

import (
	"net/http"

	"github.com/AlexKorchagin/vk-films/internal/repo"
)

type ActorFilmsHandlers struct {
	Repo *repo.Repository
}

func (a *ActorFilmsHandlers) GetAllActorFilms() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
