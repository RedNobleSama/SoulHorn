/**
* @Author: oreki
* @Date: 2021/6/5 11:30
* @Email: a912550157@gmail.com
 */

package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode string
	Address string
	HttpPort string
	JWTkey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	RedisHost   string
	RedisPort   string
)


func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误,请检查", err)
	}
	LoadServer(file)
	LoadDb(file)
}

// LoadServer 加载服务配置
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").Value()
	Address = file.Section("server").Key("Address").Value()
	HttpPort = file.Section("server").Key("HttpPort").Value()
	JWTkey = file.Section("server").Key("JWTkey").Value()
}

// LoadDb 加载数据库配置
func LoadDb(file *ini.File) {
	Db = file.Section("database").Key("Db").Value()
	DbHost = file.Section("database").Key("DbHost").Value()
	DbPort = file.Section("database").Key("DbPort").Value()
	DbUser = file.Section("database").Key("DbUser").Value()
	DbPassword = file.Section("database").Key("DbPassword").Value()
	DbName = file.Section("database").Key("DbName").Value()
}