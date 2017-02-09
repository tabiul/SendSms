package burstsms

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	maxSMSSize = 160 * 3
)

type request struct {
	PhoneNumber string `json:"phoneNumber"`
	Message     string `json:"message"`
}

type response struct {
	Responses []status `json:"responses"`
}

type status struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}

// REST contains credentials required for sending REST
type REST struct {
	sms *SMS
}

// NewREST creates a rest service to send REST
func NewREST(sms *SMS) *REST {
	return &REST{sms}
}
func validateRequest(req *request) error {
	if req.PhoneNumber == "" {
		return errors.New("Phone Number is required")
	}
	if req.Message == "" {
		return errors.New("Message is required")
	}
	if len(req.Message) > maxSMSSize {
		return fmt.Errorf("Maximum message size is %d", maxSMSSize)
	}
	return nil
}

// SendSMS is a REST service to send SMS
func (rest *REST) SendSMS(w http.ResponseWriter, r *http.Request) {
	res := response{}
	if r.Method == "POST" {
		defer r.Body.Close()
		decoder := json.NewDecoder(r.Body)
		req := request{}
		err := decoder.Decode(&req)
		if err != nil {
			res.Responses = append(res.Responses, status{false, err.Error()})
			buildResponse(http.StatusUnprocessableEntity, &res, w)
		}
		err = validateRequest(&req)
		if err != nil {
			res.Responses = append(res.Responses, status{false, err.Error()})
			buildResponse(http.StatusBadRequest, &res, w)
		}
		successChan, errorChan := rest.sms.SendSMS(req.PhoneNumber, []string{req.Message})
		for r := range successChan {
			res.Responses = append(res.Responses, status{true, fmt.Sprintf("Message send successfully to %s", r.PhoneNumber)})
		}
		for r := range errorChan {
			if r.Error.Code == InternalErrorCode {
				res.Responses = append(res.Responses, status{false, "Problem sending message due to internal server error"})
			} else {
				res.Responses = append(res.Responses, status{false, fmt.Sprintf("Problem sending message due to %s", r.Error.Description)})
			}

		}
		buildResponse(http.StatusOK, &res, w)
	} else {
		buildResponse(http.StatusMethodNotAllowed, &res, w)
	}
}

func buildResponse(httpStatus int, res *response, w http.ResponseWriter) {
	log.Printf("Response: %v", res.Responses)
	bytes, err := json.Marshal(res)
	if err != nil {
		log.Printf("Unable to build response %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(httpStatus)
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Content-Length", strconv.Itoa(len(bytes)))
		w.Write(bytes)
	}
}
