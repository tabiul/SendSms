package server

import (
	"burstsms"
	"log"
	"net/http"
)

// Run starts web server
func Run(webAppDir string) {
	bitly := burstsms.NewBitly("tabiul@gmail.com", "P@ssw0rd07c", "b2748146283bd380ac1c9fb29fe8dc4fd23ee55a", "d13ecf81986d30bf67ff73e087eb3813e54b9ab2")
	sms := burstsms.NewSMS("cddd0f7a6282d6dddbe6a3fc465b6ec4", "secret", bitly)
	handler := burstsms.NewREST(sms)
	http.HandleFunc("/sms/send", handler.SendSMS)
	http.Handle("/", http.FileServer(http.Dir(webAppDir)))
	log.Print("listening to port 8080")
	http.ListenAndServe(":8080", nil)
}
