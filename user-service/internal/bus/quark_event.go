package bus

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/go-multierror"
	"github.com/maestre3d/quark-demo/user-service/internal/event"
	"github.com/neutrinocorp/quark"
)

type QuarkEvent struct {
	broker *quark.Broker
}

func NewQuarkEvent(broker *quark.Broker) *QuarkEvent {
	return &QuarkEvent{
		broker: broker,
	}
}

func (e QuarkEvent) Publish(ctx context.Context, events ...event.Domain) error {
	errs := &multierror.Error{}

	for _, ev := range events {
		if err := e.publishEvent(ctx, ev); err != nil {
			errs = multierror.Append(err, errs)
		}
	}
	return errs.ErrorOrNil()
}

func (e QuarkEvent) publishEvent(ctx context.Context, ev event.Domain) error {
	encodedEvent, err := json.Marshal(ev)
	if err != nil {
		return err
	}

	msgType := event.NewTopicFromEvent(ev)
	msg := quark.NewMessage(e.broker.MessageIDFactory(), msgType, encodedEvent)
	msg.Subject = ev.EntityID()
	msg.Source = e.broker.BaseMessageSource + newMessageSource(ev)
	msg.ContentType = e.broker.BaseMessageContentType

	return e.broker.Publisher.Publish(ctx, msg)
}

func newMessageSource(ev event.Domain) string {
	// do not use fmt.Sprintf since it will reflect data and reduce performance
	return "/" + ev.Context() + "/" + ev.AggregateID() + "/" + ev.Entity() + "/" + ev.Action()
}
