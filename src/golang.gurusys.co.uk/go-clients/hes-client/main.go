package main

import (
    _"bytes"
    "compress/gzip"
    "errors"
        "flag"
        "fmt"
    "net/http"
    _"io/ioutil"
    "encoding/json"
        "os"
    "sort"
    "strconv"
    _"sync"
    "time"

        "golang.gurusys.co.uk/go-clients/data/hes"
        "golang.gurusys.co.uk/go-clients/getter"
)

var (
        err error
	apiHes = "https://api.gurusys.co.uk/api/v2"
	tokens = "/tokens"
        clients = "/clients"
        help = flag.Bool("help", false, "prints help info")
        login = flag.Bool("login", false, "use api_key to get & save token")
        debug = flag.Bool("debug", false, "true for debug info")
        req = flag.String("req", "clients", "get what?")
)
func main() {
        if *help {
                fmt.Println("Guru Clients ...")
                os.Exit(1)
        }

        if *login {
                fmt.Println("Logging in ...")
                os.Exit(1)
        }

        hesLive := []byte(apiHes)
        raw, err := hes.Asset("static/")
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
        doneChan := make(chan bool, 1)
        errChan := make(chan error, 1)
        respChan := make(chan []byte, 1)
        fmt.Println("Getting ...")
        switch path {
        case "client":
                url := append(hesLive, path...)
        default:
                fmt.Println("No request ...")
                os.Exit(1)
        }
        go func(url string) {
                req, err := http.NewRequest("GET", url, nil)
                if err != nil {
                        errChan <- err
                        return
                }
                hes.SetHeaders(req)
                resp, err := getter.Client.Do(req)
                if err != nil {
                        errChan <- errors.New("GET"+err.Error())
                        return
                }else {
                        if resp != nil {
                                defer resp.Body.Close()
                        }
                        cl := resp.Header.Get(getter.ContentLengthHeader)
                        icl, err := strconv.Atoi(cl)
                        if err != nil {
                                //errChan <- err
                                errChan <- errors.New("Strconv"+err.Error())
                                return
                        }
                        ubs := make([]byte, icl*3)
                        ct := resp.Header.Get(getter.ContentTypeHeader)
                        switch ct {
                                case "gzip":
                                         gzr, err := gzip.NewReader(resp.Body)
                                        if err != nil {
                                                 //errChan <- err
                                                errChan <- errors.New("gzip"+err.Error())
                                                return
                                        }
                                        defer gzr.Close()
                                        nbs, err := gzr.Read(ubs)
                                        ubs = ubs[:nbs]
                                        respChan <- ubs
                                        return
                                default:
                                        gzr, err := gzip.NewReader(resp.Body)
                                        if err != nil {
                                                //errChan <- err
                                                errChan <- errors.New("default "+err.Error())
                                                 return
                                        }
                                        defer gzr.Close()
                                        nbs, err := gzr.Read(ubs)
                                        ubs = ubs[:nbs]
                                        respChan <- ubs
                                         return

                        }
                        //respChan <- []byte(ct)
                        //return
                }
                doneChan<- true
        }(string(url))

        select {
        case <-doneChan:
                fmt.Println("Done: ", n)
        case err = <-errChan:
                fmt.Println("Error: ", err)
        case bs := <-respChan:
                go func(bs []byte) {
                        hcs := &hes.HesClients{}
                        err := json.Unmarshal(bs, hcs)
                        if err != nil {
                                fmt.Println(err)
                                errChan <- err
                                return
                        }
                }(bs)
        case <-time.After(2000 * time.Millisecond):
                fmt.Println("Timeout: ", n, " timed out")
                n++
        }

        close(doneChan)
        close(errChan)
        close(respChan)
        sort.Sort(hesclients)
        for _,strng := range strngs {
            fmt.Println(strng.String())
        }

}

