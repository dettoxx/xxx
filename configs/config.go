package configs

type Config struct {
	Environment    string
	Port           int
	AppName        string
	ServiceConfigs []ServiceConfig
}

type ServiceConfig struct {
	ServiceName string
	Url         string
	AuthKey     string
}
