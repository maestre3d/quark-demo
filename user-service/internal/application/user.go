package application

import (
	"context"

	"github.com/maestre3d/quark-demo/user-service/internal/aggregate"
	"github.com/maestre3d/quark-demo/user-service/internal/event"
	"github.com/maestre3d/quark-demo/user-service/internal/repository"
)

type User struct {
	repo     repository.User
	eventBus event.Bus
}

func NewUser(r repository.User, b event.Bus) *User {
	return &User{
		repo:     r,
		eventBus: b,
	}
}

func (u *User) Create(ctx context.Context, id, username, email string) error {
	user := aggregate.NewUser(id, username, email)
	if err := u.repo.Save(ctx, *user); err != nil {
		return err
	}
	return u.eventBus.Publish(ctx, user.PullEvents()...)
}
