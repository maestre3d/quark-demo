package event

import (
	"github.com/maestre3d/quark-demo/user-service/internal/domain"
	"github.com/neutrinocorp/quark"
)

type Domain interface {
	// Context Domain-Driven domain context or service name
	Context() string
	// Entity name
	Entity() string
	AggregateID() string
	EntityID() string
	Action() string
	Version() int
}

func NewTopicFromEvent(e Domain) string {
	return quark.FormatTopicName(domain.OrgNameAlt, e.Context(), quark.DomainEvent,
		e.Entity(), e.Action(), e.Version())
}
