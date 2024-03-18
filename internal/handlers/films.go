package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/AlexKorchagin/vk-films/gen"
	"github.com/AlexKorchagin/vk-films/internal/repo"
)

type FilmsHandlers struct {
	Repo *repo.Repository
}

func (a *FilmsHandlers) CreateFilm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		film := gen.Film{}
		err := json.NewDecoder(r.Body).Decode(&film)
		if err != nil {
			http.Error(w, "Ошибка при чтении тела запроса", http.StatusUnprocessableEntity)
			return
		}

		err = repo.SaveFilmToDatabase(film, a.Repo)
		if err != nil {
			http.Error(w, "Ошибка при сохранении данных в базу данных", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Данные успешно сохранены в базе данных"))
	}
}

func (a *FilmsHandlers) UpdateFilm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		film := gen.Film{}
		err := json.NewDecoder(r.Body).Decode(&film)
		if err != nil {
			http.Error(w, "Ошибка при чтении тела запроса", http.StatusUnprocessableEntity)
			return
		}

		filmIDStr := r.URL.Path[len("/films/"):]
		filmID, err := strconv.ParseInt(filmIDStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid film ID", http.StatusBadRequest)
			return
		}
		film.ID = filmID

		var updateQuery strings.Builder
		updateQuery.WriteString("UPDATE films SET ")

		if film.Name != nil {
			updateQuery.WriteString(fmt.Sprintf("Name = '%s', ", *film.Name))
		}
		if film.Description != nil {
			updateQuery.WriteString(fmt.Sprintf("Description = '%s', ", *film.Description))
		}
		if film.PublishDay != nil {
			updateQuery.WriteString(fmt.Sprintf("PublishDay = '%s', ", *film.PublishDay))
		}
		if film.Rating != nil {
			updateQuery.WriteString(fmt.Sprintf("Rating = %f, ", *film.Rating))
		}

		// Убираем последнюю запятую
		updateQueryStr := strings.TrimSuffix(updateQuery.String(), ", ")

		// Добавляем условие WHERE
		updateQueryStr += fmt.Sprintf(" WHERE ID = %d", film.ID)

		err = repo.UpdateFilmDB(updateQueryStr, a.Repo)
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка сохранения в базу данных: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Данные успешно сохранены в базе данных"))
	}

}

// func getFilm(w http.ResponseWriter) {
// 	film := gen.Film{}
// 	b, err := json.Marshal(film)
// 	if err != nil {
// 		fmt.Fprint(w, err)
// 	} else {
// 		fmt.Fprint(w, string(b))
// 	}

// }

func (a *FilmsHandlers) DeleteFilm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filmIDStr := r.URL.Path[len("/films/"):]
		filmID, err := strconv.ParseInt(filmIDStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid film ID", http.StatusBadRequest)
			return
		}

		err = repo.DeleteFilmByID(filmID, a.Repo)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error deleting a movie from the database: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("The movie was successfully deleted from the database"))
	}

}

func (a *FilmsHandlers) GetAllFilms() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var requestData struct {
            SortBy string `json:"sortBy"`
        }
        err := json.NewDecoder(r.Body).Decode(&requestData)
        if err != nil {
            http.Error(w, "Ошибка при чтении тела запроса", http.StatusBadRequest)
            return
        }

        sortBy := requestData.SortBy
        if sortBy == "" {
            sortBy = "Rating"
        }

        films, err := repo.GetAllFilms(sortBy, a.Repo)
        if err != nil {
            http.Error(w, "Error receiving movies", http.StatusInternalServerError)
            return
        }

 
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(films)
    }

}