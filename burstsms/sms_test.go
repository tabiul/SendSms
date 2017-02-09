package burstsms

import (
	"testing"
	"burstsms/test"
)

func TestValidNumberSendSms(t *testing.T) {
	bitly := NewBitly("tabiul@gmail.com", "P@ssw0rd07c", "b2748146283bd380ac1c9fb29fe8dc4fd23ee55a", "d13ecf81986d30bf67ff73e087eb3813e54b9ab2")
	sms := NewSMS("cddd0f7a6282d6dddbe6a3fc465b6ec4", "secret", bitly)
	successChan, errorChan := sms.SendSMS("61426686571", []string{"hello from burst sms. tabiul http://www.google.com"})
	success := []SuccessResponse{}
	error := []ErrorResponse{}
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
	bitly := NewBitly("tabiul@gmail.com", "P@ssw0rd07c", "b2748146283bd380ac1c9fb29fe8dc4fd23ee55a", "d13ecf81986d30bf67ff73e087eb3813e54b9ab2")
	sms := NewSMS("cddd0f7a6282d6dddbe6a3fc465b6ec4", "secret", bitly)
	successChan, errorChan := sms.SendSMS("123", []string{"hello from burst sms. tabiul http://www.google.com"})
	success := []SuccessResponse{}
	error := []ErrorResponse{}
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

