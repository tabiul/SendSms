package burstsms

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	bitlyShortenURL    = "https://api-ssl.bitly.com/v3/shorten"
	bitlyOAuthTokenURL = "https://api-ssl.bitly.com/oauth/access_token"
)

// Bitly is a struct that holds bitly credentials
type Bitly struct {
	clientID     string
	clientSecret string
	username     string
	password     string
}

// NewBitly creates a new bitly service
func NewBitly(username, password, clientID, clientSecret string) *Bitly {
	return &Bitly{clientID, clientSecret, username, password}
}

// ShortenURL uses bitly api to shorten a given long url
func (bitly *Bitly) ShortenURL(longURL string) (string, error) {
	accessToken, err := bitly.getAccessToken()
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf(bitlyShortenURL+"?access_token=%s&longURL=%s&format=txt", accessToken, longURL)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("Unable to shorten url %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Unable to read response body %s", err)
	}
	return string(body), nil

}

func (bitly *Bitly) getAccessToken() (string, error) {
	data := url.Values{}
	if bitly.clientID != "" && bitly.clientSecret != "" {
		data.Add("client_id", bitly.clientID)
		data.Add("client_secret", bitly.clientSecret)
	}
	request, err := http.NewRequest("POST", bitlyOAuthTokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("Error creating request %s", err)
	}
	request.SetBasicAuth(bitly.username, bitly.password)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("Error requesting for access token %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Unable to read response body %s", err)
	}
	return string(body), nil
}
