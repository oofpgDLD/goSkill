package geo

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type NodeConfig struct {
	Url         string `json:"url"`
	Password    string `json:"password"`
	MaxIdle     int    `json:"maxIdle"`
	MaxActive   int    `json:"maxActive"`
	IdleTimeout int    `json:"idleTimeout"`
}

func NewPool(conf *NodeConfig) *redis.Pool{
	return &redis.Pool{
		MaxActive:   conf.MaxActive,
		MaxIdle:     conf.MaxIdle,
		IdleTimeout: time.Duration(conf.IdleTimeout) * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(conf.Url)
			if err != nil {
				panic(err)
			}
			//验证密码
			if conf.Password != "" {
				if _, err := c.Do("AUTH", conf.Password); err != nil {
					panic(err)
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}