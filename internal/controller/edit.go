package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	files "github.com/mtricht/remedio"
	"github.com/mtricht/remedio/internal/database"
)

func EditGet(w http.ResponseWriter, r *http.Request) {
	var medication database.Medication
	database.DB.First(&medication, chi.URLParam(r, "ID"))

	tmpl, err := template.ParseFS(files.TemplateFiles, "templates/base.html", "templates/edit.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, &medication)
}

func EditPost(w http.ResponseWriter, r *http.Request) {
	var medication database.Medication
	database.DB.First(&medication, chi.URLParam(r, "ID"))
	newMedication, err := parseMedicationForm(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	database.DB.Model(medication).Updates(newMedication)
	http.Redirect(w, r, "/", http.StatusFound)
}
