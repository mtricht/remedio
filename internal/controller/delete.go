package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mtricht/remedio/internal/database"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	database.DB.Delete(&database.Medication{}, chi.URLParam(r, "ID"))
	http.Redirect(w, r, "/", http.StatusFound)
}
