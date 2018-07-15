package data_test

import (
        "bytes"
	"fmt"
        "encoding/gob"
        "github.com/gomodule/redigo/redis"
        "golang.gurusys.co.uk/data"
        //"encoding/json"
        //"testing"
)

type T struct {

        S1 string
        I1 int64
}

func ExampleGob () {

        c := data.P.Get()
        defer c.Close()
	c.Send("SELECT" , 1)

        t1 := T{"t1", 1}
        var bb bytes.Buffer
        enc := gob.NewEncoder(&bb)
        err := enc.Encode(t1)
        c.Do("SET", "t1bytes", bb.Bytes())
        rec, err := redis.Bytes(c.Do("GET", "t1bytes"))
        if err != nil {
                fmt.Println(err)
        }
        dec := gob.NewDecoder(bytes.NewBuffer(rec))
        t2 := T{}
        err = dec.Decode(&t2)
        if err != nil {
               fmt.Println("Oops", err)
        }
        fmt.Println(t2)

        // Output:
        // {t1 1}
}

