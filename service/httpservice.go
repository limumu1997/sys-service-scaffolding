package service

import (
	"encoding/json"
	"net/http"
	"strings"
	"sys-service-scaffolding/config"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func ListenAndServe() {
	r := mux.NewRouter()
	s := r.PathPrefix("/etcdfs/api").Subrouter()
	// INIT_VIUINIT_REQ_101_20230202155112230.json
	s.HandleFunc("/{gantryid}/{filetype}/{filename}", etcdfsPRAPiHandler).Methods(http.MethodPost)
	http.Handle("/", r)
	http.ListenAndServe(config.Config.ListenPort, nil)
}

// 门架系统车牌图像识别设备接口
func etcdfsPRAPiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	filename := mux.Vars(r)["filename"]
	filenametype := strings.Index(filename, "_REQ")
	filetype := filename[:filenametype]
	log.Infof("收到 %s 数据包", filename)
	switch {
	case strings.EqualFold(filetype, "INIT_VIUINIT"):
		log.Info("执行 设备初始化 数据包处理")
		initViuinitHandler(w, r)
	case strings.EqualFold(filetype, "MON_BVIUBASEINFO"):
		log.Info("执行 基础数据上传 数据包处理")
		monBviubaseinfoHandler(w, r)
	case strings.EqualFold(filetype, "TRC_BVIU"):
		log.Info("执行 图片流水上传 数据包处理")
		trcBviuHandler(w, r)
	case strings.EqualFold(filename, "TRC_BVIPU"):
		log.Info("执行 图片上传 数据包处理")
		trcBvipuHandler(w, r)
	case strings.EqualFold(filetype, "MON_BVIUSTATE"):
		log.Info("执行 状态信息上传 数据包处理")
		monBviustateHandler(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"msg" : "服务器无法理解该请求"}`))
		log.Errorf("%s 数据包无法处理", filename)
	}
}

// 设备初始化
func initViuinitHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody string
	json.NewDecoder(r.Body).Decode(&reqBody)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"gantryHex":      "2CFFFF",
		"gantryOrderNum": 1,
		"driveDir":       1,
		"receiveTime":    time.Now().Format("2006-01-02T15:04:05"),
		"subCode":        0,
		"errorMsg":       "成功",
	})
}

// 基础数据上传
func monBviubaseinfoHandler(w http.ResponseWriter, r *http.Request) {
}

// 图片流水上传
func trcBviuHandler(w http.ResponseWriter, r *http.Request) {
}

// 图片上传
func trcBvipuHandler(w http.ResponseWriter, r *http.Request) {
}

// 状态信息上传
func monBviustateHandler(w http.ResponseWriter, r *http.Request) {
}
