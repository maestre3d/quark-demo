package controller

import (
	"encoding/json"
	"log"

	"github.com/maestre3d/quark-demo/analytics-service/internal/event"
	"github.com/neutrinocorp/quark"
)

type UserPubSub struct{}

func (c *UserPubSub) SetSubscribers(b *quark.Broker) {
	b.Topic(event.NewTopicFromEvent(event.UserCreated{})).
		Group(event.NewQueueFromEvent(event.UserCreated{}, "record_user")).
		HandleFunc(c.sendEmailWhenCreated)
}

func (c *UserPubSub) sendEmailWhenCreated(w quark.EventWriter, e *quark.Event) bool {
	usrCreated := event.UserCreated{}
	if err := json.Unmarshal(e.Body.Data, &usrCreated); err != nil {
		log.Println(err)
		return false
	}
	log.Printf("new user [ID %s | Username %s] aggregated to data lake", usrCreated.ID, usrCreated.Username)
	return true
}
