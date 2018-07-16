package main

import (
        "bufio"
        _"bytes"
        _"compress/gzip"
        _"errors"
        "flag"
        "fmt"
        _"encoding/gob"
        _"net/http"
        _"io/ioutil"
        _"encoding/json"
        "os"
        _"os/signal"
        _"sort"
        _"strconv"
        _"sync"
        "time"

        _"golang.gurusys.co.uk/data/pw"
        _"golang.gurusys.co.uk/data"
        "golang.gurusys.co.uk/getter"
)

var (
        err error
	apiPw = "https://api.prosperworks.com/developer_api/v1/"
        help = flag.Bool("help", false, "prints this message")
        ping = flag.Bool("ping", false, "use access token to ping hes")
        debug = flag.Bool("debug", false, "set to true for debug info")
        path = flag.String("path", "account", "set to desired api path, eg. path=account")
)


func main() {
        interruptChan := make(chan os.Signal, 1)
        getter.SignalNotify(interruptChan)

        flag.Parse()

        cmds := make(chan string, 1)
        resp := make(chan string, 1)
        go func() {
        execute:
                a:=<-cmds
                ab := []byte(a)
                a = string(ab[0:len(ab)-1])
                if a == "exit" {
                        fmt.Printf("exiting %s\n", a)
                        resp<-"done"
                        return
                }
                fmt.Printf("executing %s\n", a)
                resp<-"done"+a
                goto execute

        }()
read:
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Cmd: ")
        cmd, _ := reader.ReadString('\n')
        cmds<-cmd
        //for {
                select {
                case r:=<-resp:
                        if r=="done" {
                                fmt.Printf("Done: %s\n", r)
                                goto end //return
                        }
                        fmt.Println(r)
                case <-interruptChan:
                        cmds<-"exit"
                        goto end
                case <-time.After(20 * time.Second):
                        cmds<-"exit"
                        fmt.Println("Timeout: timed out")
                        goto end
                }
        //}
        goto read

end:
        close(cmds)
        close(resp)
        close(interruptChan)

}
