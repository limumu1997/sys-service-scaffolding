package service

import (
	"net/http"
	"sys-service-scaffolding/config"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func ListenAndServe() {
	r := mux.NewRouter()
	r.HandleFunc("/index", IndexHandler).Methods(http.MethodPost)
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    config.Config.ListenPort,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 3 * time.Second,
		ReadTimeout:  3 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
