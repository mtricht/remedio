package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kelseyhightower/envconfig"
	files "github.com/mtricht/remedio"
	"github.com/mtricht/remedio/internal/config"
	"github.com/mtricht/remedio/internal/controller"
	"github.com/mtricht/remedio/internal/cron"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	loadConfig()
	go runCron()
	runServer()
}

func loadConfig() {
	err := envconfig.Process("remedio", &config.Config)
	if err != nil {
		log.Fatal(err)
	}
}

func runCron() {
	if config.Config.NotificationURL == "" {
		log.Fatal("Missing notification URL")
	}
	cron.RunCron()
	for range time.Tick(time.Minute) {
		cron.RunCron()
	}
}

func runServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	if config.Config.Username != "" && config.Config.Password != "" {
		r.Use(middleware.BasicAuth("remedio", map[string]string{config.Config.Username: config.Config.Password}))
	}
	r.Get("/", controller.Index)
	r.Get("/add", controller.AddGet)
	r.Post("/add", controller.AddPost)
	r.Get("/edit/{ID}", controller.EditGet)
	r.Post("/edit/{ID}", controller.EditPost)
	r.Get("/delete/{ID}", controller.Delete)
	r.Get("/take/{ID}", controller.Take)
	r.Handle("/*", http.FileServer(http.FS(files.StaticFiles)))
	log.Printf("started on port %d", config.Config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", config.Config.Port), r))
}
