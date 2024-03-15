package service

import (
	"github.com/XxThunderBlastxX/thunder-api/domain"
	"github.com/XxThunderBlastxX/thunder-api/internal/gen/smtpconfig"
	"github.com/XxThunderBlastxX/thunder-api/pkg/helper"
)

type contactMeService struct {
	smtpConfig *smtpconfig.SMTPConfig
}

func NewContactMeService(smtpConfig *smtpconfig.SMTPConfig) domain.ContactMeService {
	return &contactMeService{
		smtpConfig: smtpConfig,
	}
}

func (c *contactMeService) SendMail(msg domain.Message) error {
	if err := helper.SendMail(msg, c.smtpConfig, "koustavmondal55@gmail.com", "me@koustav.dev"); err != nil {
		return err
	}
	return nil
}
