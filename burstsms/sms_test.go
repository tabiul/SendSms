package burstsms

import (
	"github.com/tabiul/SendSms/burstsms/server"
	"github.com/tabiul/SendSms/burstsms/test"
	"os"
	"testing"
)

func TestValidNumberSendSms(t *testing.T) {
	bitlyUsername := os.Getenv(server.BitlyUsername)
	bitlyPassword := os.Getenv(server.BitlyPassword)
	bitlyClientID := os.Getenv(server.BitlyClientID)
	bitlyClientSecret := os.Getenv(server.BitlyClientSecret)
	burstsmsClientID := os.Getenv(server.BurstSMSClientID)
	burstsmsClientSecret := os.Getenv(server.BurstSMSClientSecret)

	bitly := NewBitly(bitlyUsername, bitlyPassword, bitlyClientID, bitlyClientSecret)
	sms := NewSMS(burstsmsClientID, burstsmsClientSecret, bitly)
	successChan, errorChan := sms.SendSMS("61426686571", []string{"hello from burst sms. tabiul http://www.google.com"})
	success := []*SuccessResponse{}
	error := []*ErrorResponse{}
	for r := range successChan {
		success = append(success, r)
	}
	for r := range errorChan {
		error = append(error, r)
	}
	test.AssertEquals(t, 0, len(error))
	test.AssertEquals(t, 1, len(success))
}

func TestInValidNumberSendSms(t *testing.T) {
	bitlyUsername := os.Getenv(server.BitlyUsername)
	bitlyPassword := os.Getenv(server.BitlyPassword)
	bitlyClientID := os.Getenv(server.BitlyClientID)
	bitlyClientSecret := os.Getenv(server.BitlyClientSecret)
	burstsmsClientID := os.Getenv(server.BurstSMSClientID)
	burstsmsClientSecret := os.Getenv(server.BurstSMSClientSecret)

	bitly := NewBitly(bitlyUsername, bitlyPassword, bitlyClientID, bitlyClientSecret)
	sms := NewSMS(burstsmsClientID, burstsmsClientSecret, bitly)
	successChan, errorChan := sms.SendSMS("123", []string{"hello from burst sms. tabiul http://www.google.com"})
	success := []*SuccessResponse{}
	error := []*ErrorResponse{}
	for r := range successChan {
		success = append(success, r)
	}
	for r := range errorChan {
		error = append(error, r)
	}
	test.AssertEquals(t, 1, len(error))
	test.AssertEquals(t, 0, len(success))
	test.AssertEquals(t, `Field "to" is not a valid number.`, error[0].Error.Description)
}
