package config

var App = &app{}

type app struct{}

func (a *app) AppName() string {
	return _config.Section("app").Key("app_name").String()
}

func (a *app) Host() string {
	return _config.Section("app").Key("host").String()
}

func (i *app) IsProd() bool {
	return _config.Section("app").Key("env").String() == "prod"
}

func (i *app) IsTest() bool {
	return _config.Section("app").Key("env").String() == "dev"
}

func (i *app) Port() string {
	return _config.Section("app").Key("port").String()
}
