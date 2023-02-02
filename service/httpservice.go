package service

import (
	"net/http"
	"sys-service-scaffolding/config"

	log "github.com/sirupsen/logrus"
)

func Start() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(config.Config.ListenPort, nil)
}

func Stop() {
	log.Info("disconnect")
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
	// 使用浏览器请求一次有两条日志是因为多请求一次 /favicon.ico
	log.Info(r.RemoteAddr)
}
