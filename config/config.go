package config

import (
	"embed"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

var (
	//go:embed config.json
	c      embed.FS
	Config config
)

type config struct {
	ServiceName        string `json:"service_name"`
	ServiceDisplayName string `json:"service_display_name"`
	ServiceDescription string `json:"service_description"`
	DataPath           string `json:"data_path"`
	DBIP               string `json:"db_ip"`
	ListenPort         string `json:"listen_port"`
	DialAddr           string `json:""`
}

func init() {
	executable, _ := os.Executable()
	res, _ := filepath.EvalSymlinks(filepath.Dir(executable))
	absPath := filepath.Join(res, "config.json")
	dataConfig, err := os.ReadFile(absPath)
	if err != nil {
		log.Println("prod env can not find config.json file, use embed config")
		dataConfig, _ = c.ReadFile("config.json")
	}
	json.Unmarshal(dataConfig, &Config)
}
