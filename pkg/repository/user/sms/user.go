package sms

import (
	"fmt"

	"github.com/souvikhaldar/gobudgetsms"
)

var smsConfig gobudgetsms.Details

func init() {
	smsConfig = gobudgetsms.SetConfig("jhonson", "19931", "ccf8603569cfef50cef21b31b741626b", "", 1, 0, 0)

}

func (r *sms) Sendmessage(toNumber string, otpMessage string) error {
	message := "Your OTP for xyz is " + otpMessage
	res, err := gobudgetsms.SendSMS(smsConfig, message, toNumber, "xyz")
	if err != nil {
		return err
	}
	fmt.Println("The response after sending sms is ", res)
	return nil
}
