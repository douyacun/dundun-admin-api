package config

var App = &app{}

type app struct{}

func (a *app) AppName() string {
	return _config.Section("app").Key("app_name").String()
}

func (a *app) Host() string {
	return _config.Section("app").Key("host").String()
}

func (a *app) IsProd() bool {
	return _config.Section("app").Key("env").String() == "prod"
}

func (a *app) IsTest() bool {
	return _config.Section("app").Key("env").String() == "dev"
}

func (a *app) Port() string {
	return _config.Section("app").Key("port").String()
}

func (a *app) RootPath() string {
	return _config.Section("app").Key("root_path").String()
}
