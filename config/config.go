package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	ServiceName        string
	ServiceDisplayName string
	ServiceDescription string
	DbIP               string
}

var Config config

func init() {
	executable, _ := os.Executable()
	res, _ := filepath.EvalSymlinks(filepath.Dir(executable))
	dataConfig, err := os.ReadFile(res + "/config.json")
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(dataConfig, &Config)
}
