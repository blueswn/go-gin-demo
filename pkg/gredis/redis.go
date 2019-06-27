package gredis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"

	"go-gin-test/pkg/setting"
)

var RedisConn *redis.Pool

func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:     setting.RedisSetting.MaxIdle,
		MaxActive:   setting.RedisSetting.MaxActive,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	fmt.Println("redis conn success")

	return nil
}

func SetEX(key string, value string, seconds int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value, "EX", seconds)
	if err != nil {
		return err
	}

	return nil
}

func Command(commandName string, args ...interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()

	_, err := conn.Do(commandName, args...)
	if err != nil {
		return err
	}

	return nil
}

func Exec(commandName string, args ...interface{}) ([]string,error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Strings(conn.Do(commandName, args...))
	if err != nil {
		return nil, err
	}


	return reply, nil
}

func Set(key string, value interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	//value, err := json.Marshal(data)
	//if err != nil {
	//	return err
	//}

	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func Del(args ...interface{}) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", args...))
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}

func Decr(key string) (int, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Int(conn.Do("DECR", key))
	if err != nil {
		return -1, err
	}

	return reply, nil
}

func INCR(key string) (int, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Int(conn.Do("INCR", key))
	if err != nil {
		return -1, err
	}

	return reply, nil
}

func EXPIRE(key string, sec int) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("EXPIRE", key, sec))
}

func TTL(key string) (int, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Int(conn.Do("ttl", key))
	if err != nil {
		return -3, err
	}

	return reply, nil
}