package redis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
)

func HSet(key string, field string, value interface{}) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	value, err := json.Marshal(value)
	if err != nil {
		return 0, err
	}

	i, err := redis.Int(conn.Do("HSET", key, field, value))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func HSetNx(key string, field string, value interface{}) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	value, err := json.Marshal(value)
	if err != nil {
		return 0, err
	}

	i, err := redis.Int(conn.Do("HSETNX", key, field, value))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func HGet(key string, field string) (string, error) {
	conn := pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("HGET", key, field))
	if err != nil {
		return "", err
	}

	return s, nil
}

func HDel(key string) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("HDEL", key))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func HExists(key string, field string) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("HEXISTS", key, field))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func HIncrBy(key string, field string, increment int) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("HINCRBY", key, field, increment))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func HIncrByFloat(key string, field string, increment float64) (float64, error) {
	conn := pool.Get()
	defer conn.Close()

	f, err := redis.Float64(conn.Do("HINCRBYFLOAT", key, field, increment))
	if err != nil {
		return 0.00, err
	}

	return f, nil
}

func HKeys(key string) ([]string, error) {
	conn := pool.Get()
	defer conn.Close()

	strings, err := redis.Strings(conn.Do("HKEYS", key))
	if err != nil {
		return nil, err
	}

	return strings, nil
}

func HLen(key string) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("HLEN", key))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func HStrLen(key string, field string) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("HStrLen", key, field))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func HVals(key string) ([]string, error) {
	conn := pool.Get()
	defer conn.Close()

	strings, err := redis.Strings(conn.Do("HVALS", key))
	if err != nil {
		return nil, err
	}

	return strings, nil
}
