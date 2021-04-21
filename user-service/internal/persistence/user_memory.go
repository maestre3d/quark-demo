package persistence

import (
	"context"
	"sync"

	"github.com/maestre3d/quark-demo/user-service/internal/aggregate"
)

type UserInMemory struct {
	db map[string]*aggregate.User
	mu sync.RWMutex
}

func NewUserInMemory() *UserInMemory {
	return &UserInMemory{
		db: make(map[string]*aggregate.User),
		mu: sync.RWMutex{},
	}
}

func (u *UserInMemory) Save(ctx context.Context, user aggregate.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.db[user.ID] = &user
	return nil
}

func (u *UserInMemory) Find(ctx context.Context, id string) (*aggregate.User, error) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.db[id], nil
}
