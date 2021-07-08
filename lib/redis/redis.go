package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/zewei1022/lemon-gin-web-framework/config"
	"time"
)

var pool *redis.Pool

func Initialize(config config.Redis) {
	if config.Addr != "" {
		pool = &redis.Pool{
			MaxIdle:     config.MaxIdle,
			MaxActive:   config.MaxActive,
			IdleTimeout: time.Duration(config.IdleTimeout) * time.Second,
			Wait:        config.Wait,
			Dial: func() (redis.Conn, error) {
				conn, err := redis.Dial("tcp", config.Addr)
				if err != nil {
					return nil, err
				}

				if config.Password != "" {
					if _, err := conn.Do("AUTH", config.Password); err != nil {
						_ = conn.Close()
						return nil, err
					}
				}

				if config.Db > 0 {
					if _, err := conn.Do("SELECT", config.Db); err != nil {
						_ = conn.Close()
						return nil, err
					}
				}
				return conn, nil
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		}
	}
}
