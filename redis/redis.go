package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var RedisPoolAlias map[string]*redis.Pool

const (
	LOGLEVEL       = "loglevel"
	PUBLICPASSWORD = "pw"
	LOGLEVELHOST   = "localhost:6379"
)

func Initialize() {
	RedisPoolAlias = map[string]*redis.Pool{}
	
	RedisPoolAlias[LOGLEVEL] = newRedisPool(LOGLEVELHOST, PUBLICPASSWORD, 8, 20)

}

func newRedisPool(server, pw string, db, maxActive int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		MaxActive:   maxActive,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server, redis.DialPassword(pw), redis.DialDatabase(db))
			if err != nil {
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

func GetString(targetRedis, keyword string) string {
	conn := RedisPoolAlias[targetRedis].Get()
	defer conn.Close()
	str, err := redis.String(conn.Do("GET", keyword))
	if err != nil {
		fmt.Println("GET ERROR ", err)
		return ""
	}
	return str
}

func GetList(targetRedis, key string) []string {
	conn := RedisPoolAlias[targetRedis].Get()
	defer conn.Close()
	value, err := redis.Strings(conn.Do("LRANGE", key, 0, -1))
	if err != nil {
		fmt.Println(err)
	}
	return value
}

func GetHash(targetRedis, key string) map[string]string {
	conn := RedisPoolAlias[targetRedis].Get()
	defer conn.Close()
	obj, err := redis.StringMap(conn.Do("HGETALL", key))
	if err != nil {
		fmt.Println(err)
	}
	return obj
}

func ScanKeyList(targetRedis, key string) []string {
	conn := RedisPoolAlias[targetRedis].Get()
	defer conn.Close()
	nextCursor := 0
	var result []string
	for {
		arr, err := redis.Values(conn.Do("SCAN", nextCursor, "MATCH", key, "COUNT", 1000))
		if err != nil {
			return result
		} else {
			nextCursor, _ = redis.Int(arr[0], nil)
			tempUserList, _ := redis.Strings(arr[1], nil)
			result = append(result, tempUserList...)
		}
		if nextCursor == 0 {
			break
		}
	}
	return result
}
