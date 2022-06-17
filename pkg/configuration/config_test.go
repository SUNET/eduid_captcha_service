package configuration

import (
	"eduid_captcha_service/pkg/logger"
	"eduid_captcha_service/pkg/model"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
	"gopkg.in/yaml.v3"
)

var mockConfig = []byte(`
---
eduid:
  worker:
    common:
      debug: yes
    ladok-x:
      api_server:
        host: :8080
    captcha:
      api_server:
        host: :8080
      production: false
    x_service:
      api_server:
        host: 8080
`)

func TestParse(t *testing.T) {
	tempDir := t.TempDir()

	tts := []struct {
		name           string
		setEnvVariable bool
	}{
		{
			name:           "OK",
			setEnvVariable: true,
		},
	}

	for _, tt := range tts {
		path := fmt.Sprintf("%s/test.cfg", tempDir)
		if err := os.WriteFile(path, mockConfig, 0666); err != nil {
			assert.NoError(t, err)
		}
		if tt.setEnvVariable {
			os.Setenv("EDUID_CONFIG_YAML", path)
		}

		want := &model.Config{}
		err := yaml.Unmarshal(mockConfig, want)
		assert.NoError(t, err)

		testLog := logger.Logger{
			Logger: *zaptest.NewLogger(t, zaptest.Level(zap.PanicLevel)),
		}

		t.Run(tt.name, func(t *testing.T) {
			cfg, err := Parse(&testLog)
			assert.NoError(t, err)

			assert.Equal(t, &want.EduID.Worker.Captcha, cfg)

		})
	}

}
