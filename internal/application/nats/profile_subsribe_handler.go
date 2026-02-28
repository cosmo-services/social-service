package nats

import (
	"encoding/json"
	"main/internal/domain/profile"
	"main/pkg"

	"github.com/nats-io/nats.go"
)

type ProfileSubscribeHandler struct {
	logger         pkg.Logger
	profileService *profile.ProfileService
}

func NewProfileSubscribeHandler(profileService *profile.ProfileService, logger pkg.Logger) *ProfileSubscribeHandler {
	return &ProfileSubscribeHandler{
		profileService: profileService,
		logger:         logger,
	}
}

func (p *ProfileSubscribeHandler) OnUserRegistered(msg *nats.Msg) error {
	var event UserRegistredEvent
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		p.logger.Error(err)

		return err
	}

	if err := p.profileService.CreateProfile(
		event.UserID,
		event.Username,
		event.Email,
	); err != nil {
		p.logger.Error(err)

		return err
	}

	p.logger.Infof("Profile created by user event: %s", event)

	return nil
}

func (p *ProfileSubscribeHandler) OnUserEmailChanged(msg *nats.Msg) error {
	var event UserChangeEmailEvent
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		p.logger.Error(err)

		return err
	}

	if err := p.profileService.UpdateEmail(event.UserID, event.NewEmail); err != nil {
		p.logger.Error(err)

		return err
	}

	p.logger.Infof("Profile email updated by user event: %s", event)

	return nil
}

func (p *ProfileSubscribeHandler) OnUserUsernameChanged(msg *nats.Msg) error {
	var event UserChangeUsernameEvent
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		p.logger.Error(err)

		return err
	}

	if err := p.profileService.UpdateUsername(event.UserID, event.NewUsername); err != nil {
		p.logger.Error(err)

		return err
	}

	p.logger.Infof("Profile username updated by user event: %s", event)

	return nil
}