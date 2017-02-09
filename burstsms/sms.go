package burstsms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"
)

const (
	sendSMSURL = "https://api.transmitsms.com/send-sms.json"
	//InternalErrorCode a code to indicate internal error
	InternalErrorCode = "InternalError"
)

// SuccessResponse contains details on successful sending of SMS
type SuccessResponse struct {
	MessageID   int64   `json:"message_id"`
	Recipients  int32   `json:"recipients"`
	Cost        float64 `json:"cost"`
	PhoneNumber string
}

// ErrorResponse contains error on failure to send SMS
type ErrorResponse struct {
	Error Error `json:"error"`
}

// Error contains error details
type Error struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// SMS type contains the credentials to connect to burstsms and send sms
// A instance of bitly is required for url shortening
type SMS struct {
	clientID     string
	clientSecret string
	bitly        *Bitly
}

// NewSMS creates a SMS type to allow send of SMS message
func NewSMS(clientID, clientSecret string, bitly *Bitly) *SMS {
	return &SMS{clientID, clientSecret, bitly}
}

// http://stackoverflow.com/questions/3809401/what-is-a-good-regular-expression-to-match-a-url
func (sms *SMS) findAndReplaceWithShortURL(text string) (string, error) {
	re := regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)
	unique := map[string]string{}
	for _, url := range re.FindAllString(text, -1) {
		if _, ok := unique[url]; !ok {
			unique[url] = url
		}
	}
	for k := range unique {
		shortenURL, err := sms.bitly.ShortenURL(k)
		if err != nil {
			return "", fmt.Errorf("Unable to shorten url %s", err)
		}
		text = strings.Replace(text, k, shortenURL, -1)
	}
	return text, nil
}

func createError(code string, err error) *ErrorResponse {
	return &ErrorResponse{
		Error{
			Code:        code,
			Description: err.Error(),
		},
	}
}

// SendSMS send SMS to the phoneNumber specified
func (sms *SMS) SendSMS(phoneNumber string, messages []string) (<-chan *SuccessResponse, <-chan *ErrorResponse) {
	size := len(messages)
	successChan := make(chan *SuccessResponse, size)
	errorChan := make(chan *ErrorResponse, size)
	var wg sync.WaitGroup
	wg.Add(size)
	for _, message := range messages {
		go func(message string) {
			defer wg.Done()
			messageWithShortenedURL, err := sms.findAndReplaceWithShortURL(message)
			data := url.Values{}
			data.Add("message", messageWithShortenedURL)
			data.Add("to", phoneNumber)
			request, err := http.NewRequest("POST", sendSMSURL, strings.NewReader(data.Encode()))
			if err != nil {
				errorChan <- createError(InternalErrorCode, fmt.Errorf("Error shortening url %s", err))
				return
			}
			if err != nil {
				errorChan <- createError(InternalErrorCode, fmt.Errorf("Error building post request %s", err))
				return
			}
			request.SetBasicAuth(sms.clientID, sms.clientSecret)
			request.Header.Set("message", message)
			request.Header.Set("to", phoneNumber)
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			client := &http.Client{}
			res, err := client.Do(request)
			if err != nil {
				errorChan <- createError(InternalErrorCode, fmt.Errorf("Error executing post request %s", err))
				return
			}
			defer res.Body.Close()

			bytes, err := ioutil.ReadAll(res.Body)
			if err != nil {
				errorChan <- createError(InternalErrorCode, fmt.Errorf("Error reading the response body %s", err))
				return
			}
			if res.StatusCode == http.StatusOK {
				successResponse := &SuccessResponse{}
				err = json.Unmarshal(bytes, successResponse)
				if err != nil {
					errorChan <- createError(InternalErrorCode, fmt.Errorf("Error unmarshalling success response %s", err))
					return
				}
				successResponse.PhoneNumber = phoneNumber
				successChan <- successResponse
			} else {
				errorResponse := &ErrorResponse{}
				err = json.Unmarshal(bytes, errorResponse)
				if err != nil {
					errorChan <- createError(InternalErrorCode, fmt.Errorf("Error unmarshalling error response %s", err))
					return
				}
				errorChan <- errorResponse
			}
		}(message)
	}
	go func() {
		wg.Wait()
		close(successChan)
		close(errorChan)
	}()
	return successChan, errorChan
}
