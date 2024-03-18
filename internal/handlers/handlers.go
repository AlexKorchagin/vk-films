package handlers

import (
	"net/http"

	"github.com/AlexKorchagin/vk-films/internal/repo"
)

func InitHandlers(repo *repo.Repository) {
	actorHandlers := &ActorHandlers{
		Repo: repo,
	}

	filmsHandlers := &FilmsHandlers{
		Repo: repo,
	}

	authHandlers := &AuthHandlers{
		Repo: repo,
	}

	actorFilmsHandlers := &ActorFilmsHandlers{
		Repo: repo,
	}

	filmActorsHandlers := &FilmActorsHandlers{
		Repo: repo,
	}

	http.HandleFunc("/auth", authHandlers.Refresh())
	http.HandleFunc("/auth/login", authHandlers.Login())
	http.HandleFunc("/auth/logout", authHandlers.Logout())

	http.HandleFunc("/actors", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			actorHandlers.CreateActor()(w, r)
		case http.MethodGet:
			actorHandlers.GetAllActors()(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}

	})

	http.HandleFunc("/actors/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			actorHandlers.UpdateActor()(w, r)
		case http.MethodDelete:
			actorHandlers.DeleteActor()(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}

	})
	http.HandleFunc("/actors/{id}/films", actorFilmsHandlers.GetAllActorFilms())

	http.HandleFunc("/films", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			filmsHandlers.CreateFilm()(w, r)
		case http.MethodGet:
			filmsHandlers.GetAllFilms()(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}

	})

	http.HandleFunc("/films/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			filmsHandlers.UpdateFilm()(w, r)
		case http.MethodDelete:
			filmsHandlers.DeleteFilm()(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}

	})
	http.HandleFunc("/films/{id}/actors", filmActorsHandlers.GetAllFilmActors())
}
