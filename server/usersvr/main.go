package main

import (
	"fmt"
	"usersvr/config"
	"usersvr/log"
	"usersvr/middleware/db"
)

func main() {
	config.Init()
	usersvrConf := config.GetGlobalConfig()
	fmt.Println(usersvrConf.LogConfig.LogPath)
	log.InitLog()
	log.Debug("this is test")
	log.Info("this is a info test")
	db.GetDB()
}
