package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

func InitConfig()(*AllConfig,error){
	var cfg = new(AllConfig)
	//err:=ini.MapTo(cfg,"./config/config.ini")
	err:=ini.MapTo(cfg,"/www/wwwroot/clear/lnine/config/config.ini")
	if err!=nil{
		return nil,err
	}
	return cfg,nil
}

type AllConfig struct {
	MysqlConfig `ini:"lninemysql"`
	TestMysqlConfig `ini:"testMysql"`
}

type MysqlConfig struct{
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	DbName   string `ini:"dbname"`
}
type TestMysqlConfig struct{
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	DbName   string `ini:"dbname"`
}
func GetDSN(mysqlConfig *MysqlConfig) string{
	return fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=false&loc=Local",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.DbName,
		)
}

func GetTestDSN(testMysqlConfig *TestMysqlConfig) string{
	return fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=false&loc=Local",
		testMysqlConfig.User,
		testMysqlConfig.Password,
		testMysqlConfig.Host,
		testMysqlConfig.Port,
		testMysqlConfig.DbName,
	)
}