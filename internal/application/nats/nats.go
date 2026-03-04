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
	n.natsClient.Subscribe("AUTH_STREAM", "auth.user.registered", n.profileSubscribeHandler.OnUserRegistered)
	n.natsClient.Subscribe("AUTH_STREAM", "auth.user.email.changed", n.profileSubscribeHandler.OnUserEmailChanged)
	n.natsClient.Subscribe("AUTH_STREAM", "auth.user.username.changed", n.profileSubscribeHandler.OnUserUsernameChanged)
	n.natsClient.Subscribe("AUTH_STREAM", "auth.user.activated", n.profileSubscribeHandler.OnUserActivated)
	n.natsClient.Subscribe("AUTH_STREAM", "auth.user.deleted", n.profileSubscribeHandler.OnUserDeleted)
}

var Module = fx.Options(
	fx.Provide(NewNats),
	fx.Provide(NewProfileSubscribeHandler),
)
