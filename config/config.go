package config

var GlobalConfig config

type config struct {
	Server   Server `mapstructure:"server"`
	Dao      Dao    `mapstructure:"dao"`
	Consul   string `mapstructure:"consul"`
	Rabbitmq string `mapstructure:"rabbitmq"`
}

type Server struct {
	Addr   string `mapstructure:"addr"`
	Tmpdir string `mapstructure:"tmpdir"`
}

type Dao struct {
	Redis Redis `mapstructure:"redis"`
	Mysql Mysql `mapstructure:"mysql"`
	Mongo Mongo `mapstructure:"mongo"`
}

type Redis struct {
	Uri      string `mapstructure:"uri"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
}

type Mysql struct {
	Uri      string `mapstructure:"uri"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
}

type Mongo struct {
	Uri string `mapstructure:"uri"`
}
