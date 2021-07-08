package redis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
)

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

func Incr(key string) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("INCR", key))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func IncrBy(key string, increment int) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("INCRBY", key, increment))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func IncrByFloat(key string, increment float64) (float64, error) {
	conn := pool.Get()
	defer conn.Close()

	f, err := redis.Float64(conn.Do("INCRBYFLOAT", key, increment))
	if err != nil {
		return 0.00, err
	}

	return f, nil
}

func Decr(key string) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("DECR", key))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func DecrBy(key string, decrement int) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("DECRBY", key, decrement))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func SetNx(key string, seconds int, value interface{}) (string, error) {
	conn := pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("SETEX", key, seconds, value))
	if err != nil {
		return "", err
	}

	return s, err
}

func StrLen(key string) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("STRLEN", key))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func Append(key string, value interface{}) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	value, err := json.Marshal(value)
	if err != nil {
		return 0, err
	}

	i, err := redis.Int(conn.Do("APPEND", key, value))
	if err != nil {
		return 0, err
	}

	return i, nil
}