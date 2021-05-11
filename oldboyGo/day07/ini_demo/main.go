package main

import (
	"fmt"

	"github.com/GoWholeStack/oldboyGo/day07/reflect_demo"
)

// ini 配置文件解析器

// MysalConfig mysql 配置文件结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     string `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// RedisConfig 配置文件结构体
type RedisConfig struct {
	Host     string  `ini:"host"`
	Port     int     `ini:"port"`
	Password string  `ini:"password"`
	Database int     `ini:"database"`
	Test     bool    `ini:"test"`
	Float    float64 `ini:"float"`
}

// Config ...
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func main() {
	var cfg Config
	err := reflect_demo.LocadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Println("load ini failed, err:", err)
		return
	}
	fmt.Printf("%#v", cfg)
}
