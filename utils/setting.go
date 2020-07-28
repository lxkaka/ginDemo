package utils

import (
	"github.com/BurntSushi/toml"
)

type server struct {
	Port      int
	JWTSecret string `toml:"jwt_secret"`
}

type mysql struct {
	Host     string
	Port     int
	User     string
	Password string
}

type Conf struct {
	Server server
	Mysql  mysql
}

var Settings Conf

func Setup() {
	if _, err := toml.DecodeFile("./conf/settings.toml", &Settings); err != nil {
		panic(err)
	}
}
