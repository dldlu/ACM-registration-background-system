package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"strings"
)

var Conf = &Config{}

type Config struct {
	Mysql *MysqlConfig
}

type MysqlConfig struct {
	DSN        string
	DB         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
}

func Init() {
	path, _ := os.Getwd()
	path = path + "/config/config.toml"
	if _, err := toml.DecodeFile(path, &Conf); err != nil {
		fmt.Println("配置文件读取错误: ", err)
	}
	initMysqlDSN(Conf)
}

func initMysqlDSN(conf *Config) {
	user := conf.Mysql.DbUser
	passWord := conf.Mysql.DbPassWord
	host := conf.Mysql.DbHost
	port := conf.Mysql.DbPort
	dbName := conf.Mysql.DbName
	conf.Mysql.DSN = strings.Join([]string{user, ":", passWord, "@tcp(", host, ":", port, ")/", dbName, "?charset=utf8&parseTime=true&loc=Local"}, "")
}
