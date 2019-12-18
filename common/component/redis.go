package component

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"seckill/config"
)

type redisConn struct {
	Conn redis.Conn
}

var Redis redisConn
func InitRedis() {
	url := config.AppConfig.Redis.Host + ":" + config.AppConfig.Redis.Port
	conn, err := redis.Dial("tcp", url);
	if err != nil {
		panic(fmt.Errorf("load config fail:%s\n", err));
	}
	Redis.Conn = conn
}

func (conn *redisConn)SetStr(key, value string) {
	(*conn).Conn.Do("SET", key, value)
}

