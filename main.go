package main

import (
	logger "test/logger"
)

func main() {
	// log.SetFormatter(&easy.Formatter{
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// 	LogFormat:       "[%lvl%]: %time% - %msg%",
	// })
	log := logger.New("log.txt")

	log.Print("messageé")
	log.Printf("messageé %s", 4)
	log.Trace("messageé")
	log.Tracef("messageé %s", 4)
	log.Debug("messageé")
	log.Debugf("messageé %s", 4)
	log.Info("messageé")
	log.Infof("messageé %s", 4)
	log.Warning("messageé")
	log.Warningf("messageé %s", 4)
	log.Error("messageé")
	log.Errorf("messageé %s", 4)
	log.Fatal("messageé")
	log.Fatalf("messageé %s", 4)

}
