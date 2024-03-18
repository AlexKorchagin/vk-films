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

type ActorHandlers struct {
	Repo *repo.Repository
}

func (a *ActorHandlers) CreateActor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		actor := gen.Actor{}
		err := json.NewDecoder(r.Body).Decode(&actor)
		if err != nil {
			http.Error(w, "Ошибка при чтении тела запроса", http.StatusUnprocessableEntity)
			return
		}

		err = repo.SaveActorToDatabase(actor, a.Repo)
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка сохранения в базу данных: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Данные успешно сохранены в базе данных"))

	}
}



func (a *ActorHandlers) UpdateActor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		actor := gen.Actor{}
		err := json.NewDecoder(r.Body).Decode(&actor)
		if err != nil {
			http.Error(w, "Ошибка при чтении тела запроса", http.StatusUnprocessableEntity)
			return
		}

		actorIDStr := r.URL.Path[len("/actors/"):]
		actorID, err := strconv.ParseInt(actorIDStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid actor ID", http.StatusBadRequest)
			return
		}
		actor.ID = actorID

		var updateQuery strings.Builder
		updateQuery.WriteString("UPDATE actors SET ")

		if actor.Name != nil {
			updateQuery.WriteString(fmt.Sprintf("Name = '%s', ", *actor.Name))
		}
		if actor.LastName != nil {
			updateQuery.WriteString(fmt.Sprintf("LastName = '%s', ", *actor.LastName))
		}
		if actor.Gender != nil {
			updateQuery.WriteString(fmt.Sprintf("Gender = '%s', ", *actor.Gender))
		}
		if actor.DateOfBirth != nil {
			updateQuery.WriteString(fmt.Sprintf("DateOfBirth = '%s', ", *actor.DateOfBirth))
		}

		// Убираем последнюю запятую
		updateQueryStr := strings.TrimSuffix(updateQuery.String(), ", ")

		// Добавляем условие WHERE
		updateQueryStr += fmt.Sprintf(" WHERE ID = %d", actor.ID)

		err = repo.UpdateActorDB(updateQueryStr, a.Repo)
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка сохранения в базу данных: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Данные успешно сохранены в базе данных"))
	}
}


func (a *ActorHandlers) DeleteActor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		actorIDStr := r.URL.Path[len("/actors/"):]
		actorID, err := strconv.ParseInt(actorIDStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid actor ID", http.StatusBadRequest)
			return
		}

		err = repo.DeleteActorByID(actorID, a.Repo)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error deleting an actor from the database: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("The actor was successfully deleted from the database"))
	}

}

func (a *ActorHandlers) GetAllActors() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {


        actors, err := repo.GetAllActors(a.Repo)
        if err != nil {
            http.Error(w, "Error receiving actors", http.StatusInternalServerError)
            return
        }

 
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(actors)
    }

}