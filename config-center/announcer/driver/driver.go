package driver

import "io"

type Config struct {
	Eps []string
	SecretKeyring string
}

type Driver interface {
	Create(c *Config) (IManager, error)
}

type IVersion interface {
	Flag() string
}

type IManager interface {
	//获取配置文件 精确读取
	Get(ver IVersion, filename string) (map[string][]byte, error)
	//获取所有版本
	GetAllVersion(ver IVersion, filename string) (map[string][]byte, error)
	//新增配置文件 配置文件名称-环境-服务名称-版本号
	Add(ver IVersion, filename string, rdc io.Reader) error
}