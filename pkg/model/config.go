package model

// Cfg is the main configuration structure for this application
type Cfg struct {
	APIServer struct {
		Host string `yaml:"host" validate:"required"`
	} `yaml:"api_server"`

	Production bool   `yaml:"production"`
	HTTPProxy  string `yaml:"http_proxy"`

	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`

	Sunet struct {
		Auth struct {
			URL string `yaml:"url" validate:"required,url"`
		} `yaml:"auth"`
	} `yaml:"sunet"`

	Redis struct {
		DB                  int      `yaml:"db" validate:"required"`
		Addr                string   `yaml:"host" validate:"required_without_all=SentinelHosts SentinelServiceName"`
		SentinelHosts       []string `yaml:"sentinel_hosts" validate:"required_without=Addr,omitempty,min=2,max=4"`
		SentinelServiceName string   `yaml:"sentinel_service_name" validate:"required_with=SentinelHosts"`
	} `yaml:"redis"`
}

// Config represent the complete config file structure
type Config struct {
	EduID struct {
		Worker struct {
			Captcha Cfg `yaml:"captcha"`
		} `yaml:"worker"`
	} `yaml:"eduid"`
}
