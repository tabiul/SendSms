package main

import (
	"burstsms"
	"com/burstsms/bitly"
	"com/burstsms/rest"
	"com/burstsms/sms"
	"flag"
	"log"
	"net/http"
	"os"
	"burstsms/server"
)

const (
	appName = "smsApp"
)

var webAppDir string

func init() {
	log.SetFlags(log.Ltime)
	log.SetPrefix(appName + ":")
	flag.StringVar(&webAppDir, "webapp", "", "specify the path to webapp(required)")
}

func main() {
	flag.Parse()
	if webAppDir == "" {
		log.Fatal("webapp directory is required")
	}
	if _, err := os.Stat(webAppDir); err != nil {
		log.Fatalf("webapp directory %s is not valid", webAppDir)
	}

	server.Run(webAppDir)

}
