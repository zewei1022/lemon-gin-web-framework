package redis

import (
	"github.com/gomodule/redigo/redis"
)

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

func ZRangeByLex(key string, min interface{}, max interface{}, limit bool, offset int, count int) ([]string, error) {
	conn := pool.Get()
	defer conn.Close()

	if limit {
		strings, err := redis.Strings(conn.Do("ZRANGEBYLEX", key, min, max, "LIMIT", offset, count))
		if err != nil {
			return nil, err
		}
		return strings, nil
	} else {
		strings, err := redis.Strings(conn.Do("ZRANGEBYLEX", key, min, max))
		if err != nil {
			return nil, err
		}
		return strings, nil
	}
}

func ZRangeByScore(key string, min interface{}, max interface{}, withScores bool, limit bool, offset int, count int) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()

	if withScores {
		if limit {
			stringMap, err := redis.StringMap(conn.Do("ZRANGEBYSCORE", key, min, max, "WITHSCORES", "LIMIT", offset, count))
			if err != nil {
				return nil, err
			}
			return stringMap, nil
		} else {
			stringMap, err := redis.StringMap(conn.Do("ZRANGEBYSCORE", key, min, max, "WITHSCORES"))
			if err != nil {
				return nil, err
			}
			return stringMap, nil
		}
	} else {
		if limit {
			strings, err := redis.Strings(conn.Do("ZRANGEBYSCORE", key, min, max, "LIMIT", offset, count))
			if err != nil {
				return nil, err
			}
			return strings, nil
		} else {
			strings, err := redis.Strings(conn.Do("ZRANGEBYSCORE", key, min, max))
			if err != nil {
				return nil, err
			}
			return strings, nil
		}
	}
}

func ZRank(key string, member string) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("ZRANK", key, member))
	if err != nil {
		return nil, err
	}

	return i, nil
}

func ZRem(key string, member string) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("ZREM", key, member))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func ZRemRangeByLex(key string, min interface{}, max interface{}) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("ZREMRANGEBYLEX", key, min, max))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func ZRemRangeByRank(key string, start int, stop int) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("ZREMRANGEBYRANK", key, start, stop))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func ZRemRangeByScore(key string, min interface{}, max interface{}) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("ZREMRANGEBYSCORE", key, min, max))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func ZRevRange(key string, start int, stop int, withScores bool) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()

	if withScores {
		stringMap, err := redis.StringMap(conn.Do("ZREVRANGE", key, start, stop, "WITHSCORES"))
		if err != nil {
			return nil, err
		}
		return stringMap, nil
	} else {
		s, err := redis.Strings(conn.Do("ZREVRANGE", key, start, stop))
		if err != nil {
			return nil, err
		}
		return s, nil
	}
}

func ZRevRangeByLex(key string, min interface{}, max interface{}, limit bool, offset int, count int) ([]string, error) {
	conn := pool.Get()
	defer conn.Close()

	if limit {
		strings, err := redis.Strings(conn.Do("ZREVRANGEBYLEX", key, min, max, "LIMIT", offset, count))
		if err != nil {
			return nil, err
		}
		return strings, nil
	} else {
		strings, err := redis.Strings(conn.Do("ZREVRANGEBYLEX", key, min, max))
		if err != nil {
			return nil, err
		}
		return strings, nil
	}
}

func ZRevRangeByScore(key string, min interface{}, max interface{}, withScores bool, limit bool, offset int, count int) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()

	if withScores {
		if limit {
			stringMap, err := redis.StringMap(conn.Do("ZREVRANGEBYSCORE", key, min, max, "WITHSCORES", "LIMIT", offset, count))
			if err != nil {
				return nil, err
			}
			return stringMap, nil
		} else {
			stringMap, err := redis.StringMap(conn.Do("ZREVRANGEBYSCORE", key, min, max, "WITHSCORES"))
			if err != nil {
				return nil, err
			}
			return stringMap, nil
		}
	} else {
		if limit {
			strings, err := redis.Strings(conn.Do("ZREVRANGEBYSCORE", key, min, max, "LIMIT", offset, count))
			if err != nil {
				return nil, err
			}
			return strings, nil
		} else {
			strings, err := redis.Strings(conn.Do("ZREVRANGEBYSCORE", key, min, max))
			if err != nil {
				return nil, err
			}
			return strings, nil
		}
	}
}

func ZRevRank(key string, member string) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()

	i, err := redis.Int(conn.Do("ZREVRANK", key, member))
	if err != nil {
		return nil, err
	}

	return i, nil
}

func ZScore(key string, member string) (string, error) {
	conn := pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("ZSCORE", key, member))
	if err != nil {
		return "", err
	}

	return s, nil
}
