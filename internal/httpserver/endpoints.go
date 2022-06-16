package httpserver

import (
	"context"
	"eduid_captcha_service/internal/apiv1"
	"eduid_captcha_service/pkg/helpers"

	"github.com/gin-gonic/gin"
)

func (s *Service) endpointGetCaptcha(ctx context.Context, c *gin.Context) (interface{}, error) {
	request := &apiv1.RequestGetCaptcha{}

	if err := s.bindRequest(c, request); err != nil {
		return nil, err
	}

	if err := helpers.Check(request); err != nil {
		return nil, err
	}

	reply, err := s.apiv1.GetCaptcha(ctx, request)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *Service) endpointStatus(ctx context.Context, c *gin.Context) (interface{}, error) {
	reply, err := s.apiv1.Status(ctx)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
