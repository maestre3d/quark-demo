package controller

import (
	"encoding/json"
	"log"

	"github.com/maestre3d/quark-demo/notification-service/internal/event"
	"github.com/neutrinocorp/quark"
)

type UserPubSub struct{}

func (c *UserPubSub) SetSubscribers(b *quark.Broker) {
	b.Topic(event.NewTopicFromEvent(event.UserCreated{})).
		Group(event.NewQueueFromEvent(event.UserCreated{}, "send_email")).
		HandleFunc(c.sendEmailWhenCreated)
}

func (c *UserPubSub) sendEmailWhenCreated(w quark.EventWriter, e *quark.Event) bool {
	usrCreated := event.UserCreated{}
	if err := json.Unmarshal(e.Body.Data, &usrCreated); err != nil {
		log.Println(err)
		return false
	}
	log.Printf("sending email to %s for user %s with ID %s", usrCreated.Email, usrCreated.Username, usrCreated.ID)
	return true
}
