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
	ServiceName        string
	ServiceDisplayName string
	ServiceDescription string
	DataPath           string
	DBIP               string
	ListenPort         string
	DialAddr           string
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
