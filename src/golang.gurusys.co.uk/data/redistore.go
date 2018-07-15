package data

import (
        "flag"
        "fmt"
        "github.com/gomodule/redigo/redis"
        "os"
        "time"
)


func RedisConn() (c redis.Conn, err error) {

        c, err = redis.Dial("tcp", ":7777")
        if err != nil {
                fmt.Fprintf(os.Stderr, "RedisConn: %v\n", err)
        }
        return

}

func RedisAppend() {


}

var (
	P *redis.Pool
        redisServer = flag.String("redisServer", ":7777", "")
)

func init() {
        fmt.Println("Pinging redis ...")
        P = newPool(*redisServer)
	c := P.Get()
	//c.Send("SELECT" , 2)
	//c.Send("FLUSHDB")
        r, err := redis.String(c.Do("PING"))
	if r != "PONG" || err != nil {
		fmt.Fprintf(os.Stderr, "Do() = %v, %v, want %v, %v", r, err, time.Duration(-1), nil)
	}
	c.Close()
}

func newPool(addr string) *redis.Pool {
        return &redis.Pool{
                MaxIdle: 3,
                IdleTimeout: 240 * time.Second,
                Dial: func () (redis.Conn, error) { return redis.Dial("tcp", addr) },
        }
}
