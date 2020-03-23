package fcm

import (
	"log"

	"github.com/appleboy/go-fcm"
)

func (r *tambak) PushNotification(registrationToken string) {
	// Create the message to be sent.
	msg := &fcm.Message{
		To: registrationToken,
		Data: map[string]interface{}{
			"foo": "bar",
		},
		Notification: &fcm.Notification{
			Title: "tobalobs",
			Body:  "pakan",
		},
	}

	// Send the message and receive the response without retries.
	response, err := r.fcm.Send(msg)
	if err != nil {
		log.Println(err)
	}

	log.Printf("%#v\n", response)
}
