package redis

import (
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
)

func LIndex(key string) (string, error) {
	conn := pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("LINDEX", key))
	if err != nil {
		return "", err
	}

	return s, err
}

func LInsert(key string, position string, pivot interface{}, value interface{}) (int, error) {
	if position != "BEFORE" && position != "AFTER" {
		return 0, errors.New("position must be 'BEFORE' or 'AFTER'")
	}

	pivotValue, err := json.Marshal(pivot)
	if err != nil {
		return 0, err
	}

	value, err = json.Marshal(value)
	if err != nil {
		return 0, err
	}

	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("LINSERT", key, position, pivotValue, value))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func LLen(key string) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("LLEN", key))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func LPop(key string) (string, error) {
	conn := pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("LPOP", key))
	if err != nil {
		return "", err
	}

	return s, err
}

func LPush(key string, value interface{}) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("LPUSH", key, value))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func LPushX(key string, value interface{}) (int, error) {
	value, err := json.Marshal(value)
	if err != nil {
		return 0, err
	}

	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("LPUSHX", key, value))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func LRange(key string, start int, stop int) ([]string, error) {
	conn := pool.Get()
	defer conn.Close()

	strings, err := redis.Strings(conn.Do("LRANGE", key, start, stop))
	if err != nil {
		return nil, err
	}

	return strings, nil
}

func LRem(key string, count int, value interface{}) (int, error) {
	value, err := json.Marshal(value)
	if err != nil {
		return 0, err
	}

	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("LREM", key, count, value))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func LSet(key string, index int, value interface{}) (string, error) {
	value, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	conn := pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("LSET", key, index, value))
	if err != nil {
		return "", err
	}

	return s, nil
}

func LTrim(key string, start int, stop int) (string, error) {
	conn := pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("LTRIM", key, start, stop))
	if err != nil {
		return "", err
	}

	return s, nil
}

func RPop(key string) (string, error) {
	conn := pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("RPOP", key))
	if err != nil {
		return "", err
	}

	return s, err
}

func RPopLPush(source string, destination string) (string, error) {
	conn := pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("RPOPLPUSH", source, destination))
	if err != nil {
		return "", err
	}

	return s, err
}

func RPush(key string, value interface{}) (int, error) {
	value, err := json.Marshal(value)
	if err != nil {
		return 0, err
	}

	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("RPUSH", key, value))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func RPushX(key string, value interface{}) (int, error) {
	value, err := json.Marshal(value)
	if err != nil {
		return 0, err
	}

	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("RPUSHX", key, value))
	if err != nil {
		return 0, err
	}

	return i, nil
}


