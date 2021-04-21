package repository

import (
	"context"

	"github.com/maestre3d/quark-demo/user-service/internal/aggregate"
)

type User interface {
	Save(context.Context, aggregate.User) error
	Find(context.Context, string) (*aggregate.User, error)
}
