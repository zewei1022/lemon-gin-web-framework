package redis

import (
	"encoding/json"
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

func Set(key string, value interface{}, ex int) (string, error) {
	conn := pool.Get()
	defer conn.Close()

	value, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	s, err := redis.String(conn.Do("SET", key, value, "ex", ex))
	return s, err
}

func Get(key string) ([]byte, error) {
	conn := pool.Get()
	defer conn.Close()

	bytes, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func Exists(key string) bool {
	conn := pool.Get()
	defer conn.Close()

	b, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return b
}
