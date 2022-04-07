package config

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ini/ini"
)

type Server struct {
	Url        string
	Timeout    time.Duration
	MaxClients uint16
	SlaveID    uint16
}

var ServerSetting = &Server{}

var cfg *ini.File

func Setup(file string) {
	var err error
	cfg, err = ini.Load(file)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse './config.ini': %v", err)
	}

	mapTo("server", ServerSetting)

	ServerSetting.Timeout = ServerSetting.Timeout * time.Second
	log.Println(" ======== [ServerSetting] ======== ")
	log.Printf("Url        :%s", ServerSetting.Url)
	log.Printf("Timeout    :%v", ServerSetting.Timeout)
	log.Printf("MaxClients :%v", ServerSetting.MaxClients)
	log.Printf("SlaveID    :%v", ServerSetting.SlaveID)
	fmt.Println()
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
