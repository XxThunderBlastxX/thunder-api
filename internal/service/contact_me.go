package service

import (
	"github.com/XxThunderBlast/thunder-api/domain"
	"github.com/XxThunderBlast/thunder-api/internal/helper"
)

type contactMeService struct{}

func NewContactMeService() domain.ContactMeService {
	return &contactMeService{}
}

func (c *contactMeService) SendMail(msg domain.Message) error {
	if err := helper.SendMail(msg, "koustavmondal55gmail.com", "me@koustav.dev"); err != nil {
		return err
	}
	return nil
}
