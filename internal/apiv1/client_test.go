package apiv1

import (
	"eduid_captcha_service/pkg/logger"
	"eduid_captcha_service/pkg/model"
	"testing"
)

func mockClient(t *testing.T) *Client {
	c := &Client{
		config: &model.Cfg{},
		logger: &logger.Logger{},
	}

	return c
}
