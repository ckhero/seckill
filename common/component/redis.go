package component

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"seckill/config"
)

type ReidsConn struct {
	Conn redis.Conn
}

var redisConn redis.Conn;

func InitRedis() {
	url := config.AppConfig.Redis.Host + ":" + config.AppConfig.Redis.Port
	RedisConn, err := redis.Dial("tcp", url);
	if err != nil {
		panic(fmt.Errorf("load config fail:%s\n", err));
	}
	defer RedisConn.Close()
}

func SetStr(key, value string) {
	fmt.Println(redisConn)
	_, err := redisConn.conn.Do("SET", key, value)
	if (err != nil) {
		panic(fmt.Errorf("set redis fail:%s\n", err));
	}
}