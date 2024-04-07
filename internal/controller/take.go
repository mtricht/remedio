package controller

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/mtricht/remedio/internal/database"
	"gorm.io/gorm"
)

func Take(w http.ResponseWriter, r *http.Request) {
	var medication database.Medication
	database.DB.First(&medication, chi.URLParam(r, "ID"))
	database.DB.Model(&medication).Updates(map[string]interface{}{
		"supply":     gorm.Expr("supply - ?", 1),
		"last_taken": time.Now(),
	})
	http.Redirect(w, r, "/", http.StatusFound)
}
