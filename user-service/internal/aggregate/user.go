package aggregate

import (
	"time"

	"github.com/maestre3d/quark-demo/user-service/internal/event"
)

type User struct {
	ID         string
	Username   string
	Email      string
	CreateTime time.Time

	events []event.Domain
}

func NewUser(id, username, email string) *User {
	usr := &User{
		ID:         id,
		Username:   username,
		Email:      email,
		CreateTime: time.Now().UTC(),
		events:     make([]event.Domain, 0),
	}
	usr.recordEvent(event.UserCreated{
		ID:         id,
		Username:   username,
		Email:      email,
		CreateTime: usr.CreateTime,
	})
	return usr
}

func (u *User) recordEvent(e ...event.Domain) {
	u.events = append(u.events, e...)
}

func (u *User) PullEvents() []event.Domain {
	flushedEvents := u.events
	u.events = make([]event.Domain, 0)
	return flushedEvents
}
