package burstsms

import (
	"errors"
	"testing"
	"burstsms/test"
)

func TestShortenUrl(t *testing.T) {
	bitly := NewBitly("tabiul@gmail.com", "P@ssw0rd07c", "b2748146283bd380ac1c9fb29fe8dc4fd23ee55a", "d13ecf81986d30bf67ff73e087eb3813e54b9ab2")
	url, err := bitly.ShortenURL("http://google.com")
	if err != nil {
		test.Fail(t, err)
	}
	if url == "" {
		test.Fail(t, errors.New("empty shortened url"))
	}
}
