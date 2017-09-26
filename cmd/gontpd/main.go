package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/mengzhuo/gontpd"
)

var (
	fp    = flag.String("c", "config.example.yml", "Go NTP config file")
	level = flag.String("l", "info", "Log level, debug/info/warn/error")
)

func main() {
	flag.Parse()

	nilLogger := log.New(ioutil.Discard, "", log.Ldate)

	switch *level {
	case "debug":
	case "info":
		gontpd.Debug = nilLogger
	case "warn":
		gontpd.Debug = nilLogger
		gontpd.Info = nilLogger
	case "error":
		gontpd.Debug = nilLogger
		gontpd.Info = nilLogger
		gontpd.Warn = nilLogger
	}

	cfg, err := gontpd.NewConfigFromFile(*fp)
	if err != nil {
		log.Fatal(err)
	}
	service, err := gontpd.NewService(cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(service.Serve())
}