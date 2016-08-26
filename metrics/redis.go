package metrics

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func NewRedisPool(serveraddr, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     WORKER_MAX * 2,
		MaxActive:   WORKER_MAX * 2,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", serveraddr)
			if err != nil {
				return nil, err
			}
			if password == "" {
				return c, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// select data from redis
func RSelect(conn redis.Conn, key, value string) error {
	var redis_data string
	data, err := redis.Strings(conn.Do("LRANGE", key, 0, -1))
	if err != nil {
		log.Printf("get error :%v", err)
		return ""
	}
	if err != nil {
		log.Printf("get error :%v", err)
		return ""
	}

}

// get data from redis
func RGet(conn redis.Conn, key, value string) error {
	var redis_data string
	data, err := redis.Strings(conn.Do("LRANGE", key, 0, -1))
	if err != nil {
		log.Printf("get error :%v", err)
		return ""
	}
	if err != nil {
		log.Printf("get error :%v", err)
		return ""
	}
}

// store flow data into redis
func RSet(conn redis.Conn, key, value string) error {
	var redis_data string
	data, err := redis.Strings(conn.Do("SET", key, 0, -1))
	if err != nil {
		log.Printf("get error :%v", err)
		return ""
	}
	if err != nil {
		log.Printf("get error :%v", err)
		return ""
	}
}

// flush all redis data
func RFlush(conn redis.Conn, key, value string) error {
	var redis_data string
	data, err := redis.Strings(conn.Do("FLUSH", key, 0, -1))
	if err != nil {
		log.Printf("get error :%v", err)
		return ""
	}
	if err != nil {
		log.Printf("get error :%v", err)
		return ""
	}
}

// pushs all redis data
func RPush(conn redis.Conn, key, value string) error {
	var redis_data string
	data, err := redis.Strings(conn.Do("LPUSH", key, value))
	if err != nil {
		log.Printf("get error :%v", err)
		return err
	}
	if err != nil {
		log.Printf("get error :%v", err)
		return err
	}
	return nil
}
