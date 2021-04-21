package event

import "time"

type UserCreated struct {
	ID         string    `json:"user_id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	CreateTime time.Time `json:"create_time"`
}

func (u UserCreated) Context() string {
	return "user"
}

func (u UserCreated) Entity() string {
	return "user"
}

func (u UserCreated) AggregateID() string {
	return u.ID
}

func (u UserCreated) EntityID() string {
	return u.ID
}

func (u UserCreated) Action() string {
	return "created"
}

func (u UserCreated) Version() int {
	return 1
}
