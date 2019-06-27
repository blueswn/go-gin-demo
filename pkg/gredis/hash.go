package gredis

import (
	"github.com/gomodule/redigo/redis"
)

func HSET(key string, field string, value []byte) error {
	conn := RedisConn.Get()
	defer conn.Close()

	//value, err := json.Marshal(data)
	//if err != nil {
	//	return err
	//}
	_, err := conn.Do("HSET", key, field, value)
	if err != nil {
		return err
	}

	return nil;
}

func HGET(key string, field string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("HGET", key, field))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func HEXISTS(key string, field string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("HEXISTS", key, field))
	if err != nil {
		return false
	}

	return exists
}

func HDEL(key string, field string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("HDEL", key, field))
}

func HKEYS(key string) ([]interface{}, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Values(conn.Do("HKEYS", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func HLEN(key string)  (int64, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Int64(conn.Do("HLEN", key))
	if err != nil {
		return 0, err
	}

	return reply, nil
}

func HGETALL(key string) (map[string]string, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.StringMap(conn.Do("HGETALL", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

