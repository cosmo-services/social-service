package nats

import (
	"main/pkg"

	"go.uber.org/fx"
)

type Nats struct {
	natsClient              *pkg.NatsClient
	profileSubscribeHandler *ProfileSubscribeHandler
}

func NewNats(
	natsClient *pkg.NatsClient,
	profileSubscribeHandler *ProfileSubscribeHandler,
) *Nats {
	return &Nats{
		natsClient:              natsClient,
		profileSubscribeHandler: profileSubscribeHandler,
	}
}

func (n *Nats) SetupSubscribers() {
	streamName := "AUTH_STREAM"

	n.natsClient.Subscribe(streamName, "auth.user.registered", n.profileSubscribeHandler.OnUserRegistered)
}

var Module = fx.Options(
	fx.Provide(NewNats),
	fx.Provide(NewProfileSubscribeHandler),
)
