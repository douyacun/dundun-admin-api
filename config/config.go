package config

import (
	"log"
	"os"
	"strings"

	ini "gopkg.in/ini.v1"
)

var _config *ini.File
var logFd = os.Stdout

func Init(conf string) *ini.File {
	var (
		err error
	)
	if conf == "" {
		pwd, _ := os.Getwd()
		conf = pwd + "/config.ini"
	}
	_config, err = ini.Load(conf)
	if err != nil {
		log.Fatalf("config: %s load filed, %v", conf, err)
		return nil
	}
	return _config
}

func GetKey(key string) *ini.Key {
	parts := strings.Split(key, "::")
	section := parts[0]
	keyStr := parts[1]
	return _config.Section(section).Key(keyStr)
}

func GetLogFD() *os.File {
	return logFd
}
