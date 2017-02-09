package server

import (
	"github.com/tabiul/SendSms/burstsms"
	"log"
	"net/http"
	"os"
)

const (
	// BitlyUsername is the username to connect to bitly
	BitlyUsername        = "BITLY_USERNAME"
	// BitlyPassword is the password to connect to bitly
	BitlyPassword        = "BITLY_PASSWORD"
	// BitlyClientID is the client ID of the application registered in bitly
	BitlyClientID        = "BITLY_CLIENTID"
	// BitlyClientSecret is the client secret of the application registered in bitly
	BitlyClientSecret    = "BITLY_CLIENTSECRET"
	// BurstSMSClientID is the client ID to connect to burstsms
	BurstSMSClientID     = "BURSTSMS_CLIENTID"
	// BurstSMSClientSecret is the client secret to connect to burst sms
	BurstSMSClientSecret = "BURSTSMS_CLIENTSECRET"
)

// Run starts web server
func Run(webAppDir string) {
	bitlyUsername := os.Getenv(BitlyUsername)
	if bitlyUsername == "" {
		log.Fatalf("Environment variable %s is required", BitlyUsername)
	}

	bitlyPassword := os.Getenv(BitlyPassword)
	if bitlyPassword == "" {
		log.Fatalf("Environment variable %s is required", BitlyPassword)
	}

	burstSMSClientID := os.Getenv(BurstSMSClientID)
	if burstSMSClientID == "" {
		log.Fatalf("Environment variable %s is required", BurstSMSClientID)
	}

	burstSMSClientSecret := os.Getenv(BurstSMSClientSecret)
	if burstSMSClientSecret == "" {
		log.Fatalf("Environment variable %s is required", BurstSMSClientSecret)
	}

	bitlyClientID := os.Getenv(BitlyClientID)
	bitlyClientSecret := os.Getenv(BitlyClientSecret)

	bitly := burstsms.NewBitly(bitlyUsername, bitlyPassword, bitlyClientID, bitlyClientSecret)
	sms := burstsms.NewSMS(burstSMSClientID, burstSMSClientSecret, bitly)
	handler := burstsms.NewREST(sms)
	http.HandleFunc("/sms/send", handler.SendSMS)
	http.Handle("/", http.FileServer(http.Dir(webAppDir)))
	log.Print("listening to port 8080")
	http.ListenAndServe(":8080", nil)
}
