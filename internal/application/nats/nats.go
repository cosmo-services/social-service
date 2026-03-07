package nats

import (
	"main/internal/domain"
	"main/pkg"

	"go.uber.org/fx"
)

type Nats struct {
	natsClient              *pkg.NatsClient
	eventBus                *domain.EventBus
	profileEventHandler     *ProfileEventHandler
	profileSubscribeHandler *ProfileSubscribeHandler
}

func NewNats(
	natsClient *pkg.NatsClient,
	eventBus *domain.EventBus,
	profileEventHandler *ProfileEventHandler,
	profileSubscribeHandler *ProfileSubscribeHandler,
) *Nats {
	return &Nats{
		natsClient:              natsClient,
		eventBus:                eventBus,
		profileEventHandler:     profileEventHandler,
		profileSubscribeHandler: profileSubscribeHandler,
	}
}

func (n *Nats) SetupSubscribers() {
	n.natsClient.Subscribe("AUTH_STREAM", "auth.user.registered", n.profileSubscribeHandler.OnUserRegistered)
	n.natsClient.Subscribe("AUTH_STREAM", "auth.user.email.changed", n.profileSubscribeHandler.OnUserEmailChanged)
	n.natsClient.Subscribe("AUTH_STREAM", "auth.user.username.changed", n.profileSubscribeHandler.OnUserUsernameChanged)
	n.natsClient.Subscribe("AUTH_STREAM", "auth.user.activated", n.profileSubscribeHandler.OnUserActivated)
	n.natsClient.Subscribe("AUTH_STREAM", "auth.user.deleted", n.profileSubscribeHandler.OnUserDeleted)

	n.natsClient.Subscribe("FILE_STREAM", "file.avatar.uploaded", n.profileSubscribeHandler.OnAvatarUploaded)
	n.natsClient.Subscribe("FILE_STREAM", "file.deleted", n.profileSubscribeHandler.OnAvatarDeleted)
}

func (n *Nats) SetupPublishers() {
	n.eventBus.On("avatar.changed", n.profileEventHandler.AvatarChanged)
	n.eventBus.On("avatar.orphaned", n.profileEventHandler.AvatarOrphaned)
}

var Module = fx.Options(
	fx.Provide(NewNats),
	fx.Provide(NewProfileSubscribeHandler),
	fx.Provide(NewProfileEventHandler),
)
