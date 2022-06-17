package model

// Cfg is the main configuration structure for this application
type Cfg struct {
	APIServer struct {
		Host string `yaml:"host" validate:"required"`
	} `yaml:"api_server"`

	Production bool   `yaml:"production"`
	HTTPProxy  string `yaml:"http_proxy"`
}

// Config represent the complete config file structure
type Config struct {
	EduID struct {
		Worker struct {
			Captcha Cfg `yaml:"captcha"`
		} `yaml:"worker"`
	} `yaml:"eduid"`
}
