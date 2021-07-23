package redis

import "github.com/gomodule/redigo/redis"

func ZAdd(key string, score int, member string) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("ZADD", key, score, member))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func ZCard(key string) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("ZCARD", key))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func ZCount(key string, min interface{}, max interface{}) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("ZCOUNT", key, min, max))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func ZIncrBy(key string, increment int, member string) (string, error) {
	conn := pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("ZINCRBY", key, increment, member))
	if err != nil {
		return "", err
	}

	return s, nil
}

func ZPopMax(key string, count int) (map[string]string, error) {
	conn := pool.Get()
	defer conn.Close()

	stringMap, err := redis.StringMap(conn.Do("ZPOPMAX", key, count))
	if err != nil {
		return nil, err
	}

	return stringMap, nil
}

func ZPopMin(key string, count int) (map[string]string, error) {
	conn := pool.Get()
	defer conn.Close()

	stringMap, err := redis.StringMap(conn.Do("ZPOPMIN", key, count))
	if err != nil {
		return nil, err
	}

	return stringMap, nil
}

func ZRange(key string, start int, stop int, withScores bool) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()

	if withScores {
		stringMap, err := redis.StringMap(conn.Do("ZRANGE", key, start, stop, "WITHSCORES"))
		if err != nil {
			return nil, err
		}
		return stringMap, nil
	} else {
		s, err := redis.Strings(conn.Do("ZRANGE", key, start, stop))
		if err != nil {
			return nil, err
		}
		return s, nil
	}
}