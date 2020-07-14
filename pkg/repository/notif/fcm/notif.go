package fcm

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/messaging"

	"github.com/ws-tobalobs/pkg/models"
)

func (r *FCM) PushNotification(deviceID []string, msg models.MessagePushNotif) {
	// Create the message to be sent.
	message := &messaging.MulticastMessage{
		Data: map[string]string{
			"ID": msg.ID,
		},
		Notification: &messaging.Notification{
			Title: msg.Title,
			Body:  msg.Body,
		},
		Tokens: deviceID,
	}

	br, err := r.fcm.SendMulticast(context.Background(), message)
	if err != nil {
		log.Fatalln(err)
	}
	// log.Println("succes : ", br.SuccessCount)
	if br.FailureCount > 0 {
		var failedTokens []string
		for idx, resp := range br.Responses {
			if !resp.Success {
				// The order of responses corresponds to the order of the registration tokens.
				failedTokens = append(failedTokens, deviceID[idx])
			}
		}

		fmt.Printf("List of tokens that caused failures: %v\n", failedTokens)
	}
}
