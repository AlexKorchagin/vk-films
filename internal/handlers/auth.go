package handlers

import (
	"net/http"

	"github.com/AlexKorchagin/vk-films/internal/repo"
)

type AuthHandlers struct {
	Repo *repo.Repository
}

func (a *AuthHandlers) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (a *AuthHandlers) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (a *AuthHandlers) Refresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
