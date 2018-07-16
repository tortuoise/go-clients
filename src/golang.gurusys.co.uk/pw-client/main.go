package main

import (
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

        for {
                select {
                case <-interruptChan:
                        goto end
                case <-time.After(5000 * time.Millisecond):
                        fmt.Println("Timeout: timed out")
                        goto end
                }
        }

end:

}
