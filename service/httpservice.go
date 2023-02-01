package service

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func Start() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}

func Stop() {
	log.Info("disconnect")
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
	log.Info(time.Now())
	log.Error(time.Now().String() + "Error")
}
