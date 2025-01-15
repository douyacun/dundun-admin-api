package config

var Client = &client{}

type client struct{}

func (d *client) MysqlDsn() string {
	return _config.Section("client").Key("mysql").String()
}

func (d *client) RedisAddr() string {
	return _config.Section("client").Key("redis_addr").String()
}

func (d *client) RedisPassword() string {
	return _config.Section("client").Key("redis_password").String()
}

func (d *client) RedisDB() int {
	db, _ := _config.Section("client").Key("redis_db").Int()
	return db
}
