package nats

import (
	"context"
	"main/internal/domain"
	"main/pkg"
	"time"
)

type ProfileEventHandler struct {
	natsClient *pkg.NatsClient
	logger     pkg.Logger
}

func NewProfileEventHandler(
	natsClient *pkg.NatsClient,
	logger pkg.Logger,
) *ProfileEventHandler {
	return &ProfileEventHandler{
		natsClient: natsClient,
		logger:     logger,
	}
}

func (p *ProfileEventHandler) AvatarChanged(event domain.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := p.natsClient.PublishJSON(ctx, "profile.avatar.changed", event)
	if err != nil {
		p.logger.Error(err)

		return err
	}

	p.logger.Infof("Profile event successfully delivered: %s", event)

	return nil
}

func (p *ProfileEventHandler) AvatarOrphaned(event domain.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := p.natsClient.PublishJSON(ctx, "file.orphaned", event)
	if err != nil {
		p.logger.Error(err)

		return err
	}

	p.logger.Infof("Profile event successfully delivered: %s", event)

	return nil
}
