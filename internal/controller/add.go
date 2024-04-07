package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	files "github.com/mtricht/remedio"
	"github.com/mtricht/remedio/internal/database"
)

func AddGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(files.TemplateFiles, "templates/base.html", "templates/add.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	medication, err := parseMedicationForm(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	db.Create(medication)
	http.Redirect(w, r, "/", http.StatusFound)
}

func parseMedicationForm(r *http.Request) (*database.Medication, error) {
	r.ParseForm()
	time, err := time.Parse("15:04", r.FormValue("time"))
	if err != nil {
		return nil, err
	}
	supply, err := strconv.Atoi(r.FormValue("supply"))
	if err != nil {
		return nil, err
	}
	return &database.Medication{
		Name:   r.FormValue("name"),
		Time:   time,
		Supply: supply,
	}, nil
}
