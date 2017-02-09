package main

import (
	"flag"
	"github.com/tabiul/SendSms/burstsms/server"
	"log"
	"os"
)

const (
	appName = "smsApp"
)

var webAppDir string

func init() {
	log.SetFlags(log.Ltime)
	log.SetPrefix(appName + ":")
	flag.StringVar(&webAppDir, "webapp", "", "path to webapp(required)")
}

func main() {
	flag.Parse()
	if webAppDir == "" {
		log.Fatal("Please specify the path to webapp")
	}
	if _, err := os.Stat(webAppDir); err != nil {
		log.Fatalf("Directory %s is not valid", webAppDir)
	}

	server.Run(webAppDir)

}
