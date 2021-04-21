package app

import (
	"time"

	"github.com/Shopify/sarama"
	"github.com/maestre3d/quark-demo/analytics-service/pkg/controller"
	"github.com/neutrinocorp/quark"
	"github.com/neutrinocorp/quark/bus/kafka"
)

func NewSubscriber() *quark.Broker {
	b := newQuarkBroker()

	usrPubSub := &controller.UserPubSub{}
	usrPubSub.SetSubscribers(b)

	return b
}

func newQuarkBroker() *quark.Broker {
	broker := kafka.NewKafkaBroker(newSaramaConfig(),
		quark.WithCluster("localhost:19092", "localhost:29092", "localhost:39092"),
		quark.WithBaseMessageSource("https://cosmos.neutrinocorp.org/analytics"),
		quark.WithBaseMessageContentType("application/cloudevents+json"))
	return broker
}

func newSaramaConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.ClientID = "analytics-service"
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Return.Errors = true
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Consumer.Retry.Backoff = time.Second * 10
	return config
}
