package burstsms

import (
	"errors"
	"github.com/tabiul/SendSms/burstsms/server"
	"github.com/tabiul/SendSms/burstsms/test"
	"os"
	"testing"
)

func TestShortenUrl(t *testing.T) {
	bitlyUsername := os.Getenv(server.BitlyUsername)
	bitlyPassword := os.Getenv(server.BitlyPassword)
	bitlyClientID := os.Getenv(server.BitlyClientID)
	bitlyClientSecret := os.Getenv(server.BitlyClientSecret)

	bitly := NewBitly(bitlyUsername, bitlyPassword, bitlyClientID, bitlyClientSecret)
	url, err := bitly.ShortenURL("http://google.com")
	if err != nil {
		test.Fail(t, err)
	}
	if url == "" {
		test.Fail(t, errors.New("empty shortened url"))
	}
}
