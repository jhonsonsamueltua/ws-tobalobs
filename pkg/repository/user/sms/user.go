package sms

import (
	"fmt"
	"net/http"

	// "github.com/souvikhaldar/gobudgetsms"
	"github.com/nexmo-community/nexmo-go"
)

// var smsConfig gobudgetsms.Details

func init() {
	// smsConfig = gobudgetsms.SetConfig("jhonson", "19931", "ccf8603569cfef50cef21b31b741626b", "", 1, 0, 0)

}

func (r *sms) Sendmessage(toNumber string, otpMessage string) error {
	message := "Tobalobs kode: " + otpMessage + ". Berlaku untuk 5 menit. "
	// res, err := gobudgetsms.SendSMS(smsConfig, message, toNumber, "xyz")
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("The response after sending sms is ", res)
	// return nil
	// Auth
	auth := nexmo.NewAuthSet()
	auth.SetAPISecret("182d8d9e", "1lDMGPIOHRWtTOnE")

	// Init Nexmo
	client := nexmo.NewClient(http.DefaultClient, auth)

	// SMS
	smsContent := nexmo.SendSMSRequest{
		From: "tobalobs",
		To:   toNumber,
		Text: message,
	}

	smsResponse, _, err := client.SMS.SendSMS(smsContent)

	if err != nil {
		return err
	}

	fmt.Println("Status:", smsResponse.Messages[0].Status)
	return nil
}
