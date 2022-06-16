package httpserver

import (
	"context"
	"eduid_captcha_service/internal/apiv1"
	"eduid_captcha_service/pkg/model"
)

// Apiv1 interface
type Apiv1 interface {
	GetCaptcha(ctx context.Context, indata *apiv1.RequestGetCaptcha) (*apiv1.ReplyGetCaptcha, error)
	Status(ctx context.Context) (*model.Status, error)
}
