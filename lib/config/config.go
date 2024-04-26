package Config

import (
	"flag"
)

type Config struct {
	Addr string
}

var confTemp *Config

func GetConfig() *Config {
	if confTemp != nil {
		return confTemp
	}
	c := &Config{}
	// 使用 flag 包来定义命令行参数
	flag.StringVar(&c.Addr, "addr", ":4396", "地址")
	// 解析命令行参数
	flag.Parse()
	confTemp = c
	return c
}
