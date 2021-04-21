package event

import (
	"context"
)

type Bus interface {
	Publish(context.Context, ...Domain) error
}
