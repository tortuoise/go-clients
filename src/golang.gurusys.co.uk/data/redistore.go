package data

import (
        "fmt"
        "github.com/gomodule/redigo/redis"
        "os"
)


func RedisConn() (c redis.Conn, err error) {

        c, err = redis.Dial("tcp", ":7777")
        if err != nil {
                fmt.Fprintf(os.Stderr, "RedisConn: %v\n", err)
        }
        return

}
