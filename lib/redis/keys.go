package redis

import "github.com/gomodule/redigo/redis"

func Exists(key string) (bool, error) {
	conn := pool.Get()
	defer conn.Close()

	b, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false, err
	}

	return b, nil
}

func Expire(key string, ex int) (bool, error) {
	conn := pool.Get()
	defer conn.Close()

	b, err := redis.Bool(conn.Do("EXPIRE", key, ex))
	if err != nil {
		return false, err
	}

	return b, nil
}

func Ttl(key string) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("TTL", key))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func Del(key ...string) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	delCount := 0

	for _, value := range key {
		i, err := redis.Int(conn.Do("DEL", value))
		if err != nil {
			return delCount, err
		}
		delCount += i
	}

	return delCount, nil
}

func Rename(key string, newKey string) (string, error) {
	conn := pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("RENAME", key, newKey))
	if err != nil {
		return "", err
	}

	return s, nil
}

func RenameNx(key string, newKey string) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("RENAMENX", key, newKey))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func Keys(pattern string) ([]string, error) {
	conn := pool.Get()
	defer conn.Close()

	strings, err := redis.Strings(conn.Do("KEYS", pattern))
	if err != nil {
		return nil, err
	}

	return strings, nil
}

func Persist(key string) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("PERSIST", key))
	if err != nil {
		return 0, err
	}

	return i, nil
}