package config_center

import "time"

type Config struct {
	srvName string
	cfgName string
	ver string
	updateTime time.Time
	createTime time.Time
}

func NewConfig(b []byte) *Config{
	c := &Config{}
	//c.ParseName()
	return c
}


func (c *Config) ParseName(s string) error{
	return nil
}


func (c *Config) GetName() string{
	return ""
}