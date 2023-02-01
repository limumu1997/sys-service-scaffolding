package main

import (
	"flag"
	"os"
	"sys-service-scaffolding/config"
	httpservice "sys-service-scaffolding/service"
	_ "sys-service-scaffolding/util"

	"github.com/kardianos/service"
	log "github.com/sirupsen/logrus"
)

// service manager
var serviceType = flag.String("s", "", "Services Management, install, uninstall")

type serviceConfig struct {
	Name        string
	DisplayName string
	Description string
	Version     string
}

var ServiceConfig = &serviceConfig{
	Name:        config.Config.ServiceName,
	DisplayName: config.Config.ServiceDisplayName,
	Description: config.Config.ServiceDescription,
	Version:     "1.0.0"}

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *program) run() {
	// run anything...
	httpservice.Start()
}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func getService() service.Service {
	options := make(service.KeyValue)
	svcConfig := &service.Config{
		Name:        ServiceConfig.Name,
		DisplayName: ServiceConfig.DisplayName,
		Description: ServiceConfig.Description,
		Option:      options,
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatalln(err)
	}
	return s
}

func installService() {
	s := getService()
	status, err := s.Status()
	if err != nil && status == service.StatusUnknown {
		// 服务未知，创建服务
		if err = s.Install(); err == nil {
			s.Start()
			log.Printf("install %s install successful!", ServiceConfig.Name)
			return
		}
		log.Printf("install %s service failure, ERR: %s\n", err, ServiceConfig.Name)
	}

	if status != service.StatusUnknown {
		log.Printf("%s service installed, no reinstallation required", ServiceConfig.Name)
	}
}

func uninstallService() {
	s := getService()
	s.Stop()
	os.Exit(1)
	if err := s.Uninstall(); err == nil {
		log.Printf("%s service uninstall successful!", ServiceConfig.Name)
	} else {
		log.Printf("%s service uninstall failure, ERR: %s\n", err, ServiceConfig.Name)
	}
}

func main() {
	flag.Parse()
	switch *serviceType {
	case "install":
		installService()
	case "uninstall":
		uninstallService()
	default:
		s := getService()
		status, _ := s.Status()
		if status != service.StatusUnknown {
			// service runs
			s.Run()
		} else {
			log.Println("non-service runs")
			switch s.Platform() {
			case "windows-service":
				log.Printf("service runs: .\\%s.exe -s install", ServiceConfig.Name)
			default:
				log.Printf("service runs: sudo ./%s -s install", ServiceConfig.Name)
			}
			// run anything
			s.Run()
		}
	}
}
