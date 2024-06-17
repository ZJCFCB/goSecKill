package main

import (
	"goSecKill/pkg/mysql"

	"goSecKill/pkg/bootstrap"
	conf "goSecKill/pkg/config"
	"goSecKill/sk-app/setup"
)

func main() {
	mysql.InitMysql(conf.MysqlConfig.Host, conf.MysqlConfig.Port, conf.MysqlConfig.User, conf.MysqlConfig.Pwd, conf.MysqlConfig.Db) // conf.MysqlConfig.Db
	//setup.InitEtcd()
	setup.InitZk()
	setup.InitServer(bootstrap.HttpConfig.Host, bootstrap.HttpConfig.Port)

}
