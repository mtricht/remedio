package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	files "github.com/mtricht/remedio"
	"github.com/mtricht/remedio/internal/database"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 not found")
		return
	}
	var medication []database.Medication
	database.DB.Find(&medication)
	tmpl, err := template.ParseFS(files.TemplateFiles, "templates/base.html", "templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, &medication)
}
