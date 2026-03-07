package nats

import (
	"encoding/json"
	"main/internal/config"
	"main/internal/domain/profile"
	"main/pkg"

	"github.com/nats-io/nats.go"
)

type ProfileSubscribeHandler struct {
	logger         pkg.Logger
	profileService *profile.ProfileService

	fileUrl string
}

func NewProfileSubscribeHandler(profileService *profile.ProfileService, logger pkg.Logger, env config.Env) *ProfileSubscribeHandler {
	return &ProfileSubscribeHandler{
		profileService: profileService,
		logger:         logger,

		fileUrl: env.FileUrl,
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

func (p *ProfileSubscribeHandler) OnUserActivated(msg *nats.Msg) error {
	var event UserActivateEvent
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		p.logger.Error(err)

		return err
	}

	if err := p.profileService.Activate(event.UserID); err != nil {
		p.logger.Error(err)

		return err
	}

	p.logger.Infof("Profile activated by user event: %s", event)

	return nil
}

func (p *ProfileSubscribeHandler) OnUserDeactivated(msg *nats.Msg) error {
	var event UserDeactivateEvent
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		p.logger.Error(err)

		return err
	}

	if err := p.profileService.Deactivate(event.UserID); err != nil {
		p.logger.Error(err)

		return err
	}

	p.logger.Infof("Profile deactivated by user event: %s", event)

	return nil
}

func (p *ProfileSubscribeHandler) OnUserDeleted(msg *nats.Msg) error {
	var event UserDeleteEvent
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		p.logger.Error(err)

		return err
	}

	if err := p.profileService.DeleteProfile(event.UserID); err != nil {
		p.logger.Error(err)

		return err
	}

	p.logger.Infof("Profile deleted by user event: %s", event)

	return nil
}

func (p *ProfileSubscribeHandler) OnAvatarUploaded(msg *nats.Msg) error {
	var event AvatarUploadedEvent
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		p.logger.Error(err)

		return err
	}

	filePath := p.fileUrl + event.Directory + "/" + event.FileName
	if err := p.profileService.ChangeAvatar(event.UserID, filePath); err != nil {
		p.logger.Error(err)

		return err
	}

	p.logger.Infof("Profile avatar updated by file event: %s", event)

	return nil
}

func (p *ProfileSubscribeHandler) OnAvatarDeleted(msg *nats.Msg) error {
	var event UserFileDeletedEvent
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		p.logger.Error(err)

		return err
	}

	if event.Directory != "avatar" {
		return nil
	}

	filePath := p.fileUrl + event.Directory + "/" + event.FileName
	if err := p.profileService.ChangeAvatar(event.UserID, filePath); err != nil {
		p.logger.Error(err)

		return err
	}

	p.logger.Infof("Profile avatar deleted by file event: %s", event)

	return nil
}
