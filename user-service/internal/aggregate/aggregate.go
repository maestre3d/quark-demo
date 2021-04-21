package aggregate

import "github.com/maestre3d/quark-demo/user-service/internal/event"

type Aggregate interface {
	recordEvent(...event.Domain)
	PullEvents() []event.Domain
}
