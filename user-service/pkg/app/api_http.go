package app

import (
	"net/http"
	"time"

	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
	"github.com/maestre3d/quark-demo/user-service/internal/application"
	"github.com/maestre3d/quark-demo/user-service/internal/bus"
	"github.com/maestre3d/quark-demo/user-service/internal/persistence"
	"github.com/maestre3d/quark-demo/user-service/pkg/controller"
	"github.com/neutrinocorp/quark"
	"github.com/neutrinocorp/quark/bus/kafka"
)

func NewHTTPAPI() *http.Server {
	usrRepo := persistence.NewUserInMemory()
	broker := newQuarkBroker()
	eventBus := bus.NewQuarkEvent(broker)

	usrApp := application.NewUser(usrRepo, eventBus)

	usrCtrl := controller.UserHTTP{App: usrApp}

	router := mux.NewRouter()

	usrCtrl.SetEndpoints(router)

	return &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
}

func newQuarkBroker() *quark.Broker {
	broker := kafka.NewKafkaBroker(newSaramaConfig(),
		quark.WithCluster("localhost:19092", "localhost:29092", "localhost:39092"),
		quark.WithBaseMessageSource("https://cosmos.neutrinocorp.org"),
		quark.WithBaseMessageContentType("application/cloudevents+json"))
	return broker
}

func newSaramaConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.ClientID = "user-service"
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Return.Errors = true
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Consumer.Retry.Backoff = time.Second * 10
	return config
}
