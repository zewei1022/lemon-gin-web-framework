package redis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
)

func SAdd(key string, member ...interface{}) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	value, err := json.Marshal(member)
	if err != nil {
		return 0, err
	}

	i, err := redis.Int(conn.Do("SADD", key, value))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func SCard(key string) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("SADD", key))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func SDiff(key ...string) ([]string, error) {
	conn := pool.Get()
	defer conn.Close()

	strings, err := redis.Strings(conn.Do("SDIFF", key))
	if err != nil {
		return nil, err
	}

	return strings, nil
}

func SDiffStore(destination string, key ...string) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("SDIFFSTORE", destination, key))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func SInter(key ...string) ([]string, error) {
	conn := pool.Get()
	defer conn.Close()

	strings, err := redis.Strings(conn.Do("SINTER", key))
	if err != nil {
		return nil, err
	}

	return strings, nil
}

func SInterStore(destination string, key ...string) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("SINTERSTORE", destination, key))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func SIsMember(key string, member interface{}) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	value, err := json.Marshal(member)
	if err != nil {
		return 0, err
	}

	i, err := redis.Int(conn.Do("SISMEMBER", key, value))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func SMembers(key string) ([]string, error) {
	conn := pool.Get()
	defer conn.Close()

	strings, err := redis.Strings(conn.Do("SMEMBERS", key))
	if err != nil {
		return nil, err
	}

	return strings, nil
}

func SMove(source string, destination string, member interface{}) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	value, err := json.Marshal(member)
	if err != nil {
		return 0, err
	}

	i, err := redis.Int(conn.Do("SMOVE", source, destination, value))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func SPop(key string, count int) ([]string, error) {
	conn := pool.Get()
	defer conn.Close()

	strings, err := redis.Strings(conn.Do("SPOP", key, count))
	if err != nil {
		return nil, err
	}

	return strings, nil
}

func SRandMember(key string, count int) ([]string, error) {
	conn := pool.Get()
	defer conn.Close()

	strings, err := redis.Strings(conn.Do("SRANDMEMBER", key, count))
	if err != nil {
		return nil, err
	}

	return strings, nil
}

func SRem(key string, member ...interface{}) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	value, err := json.Marshal(member)
	if err != nil {
		return 0, err
	}

	i, err := redis.Int(conn.Do("SREM", key, value))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func SUnion(key ...string) ([]string, error) {
	conn := pool.Get()
	defer conn.Close()

	strings, err := redis.Strings(conn.Do("SUNION", key))
	if err != nil {
		return nil, err
	}

	return strings, nil
}

func SUnionStore(destination string, key ...string) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("SUNIONSTORE", destination, key))
	if err != nil {
		return 0, err
	}

	return i, nil
}