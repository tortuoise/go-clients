package main

import (
        "bytes"
        "compress/gzip"
        "errors"
        "flag"
        "fmt"
        _"encoding/gob"
        "net/http"
        "io/ioutil"
        "encoding/json"
        "os"
        "sort"
        "strconv"
        _"sync"
        "time"

        "golang.gurusys.co.uk/data/hes"
        "golang.gurusys.co.uk/data"
        "golang.gurusys.co.uk/getter"
)

var (
        err error
        hestkn = "src/golang.gurusys.co.uk/data/hes/static/accessToken"
	apiHes = "https://api.gurusys.co.uk/api/v2/"
        help = flag.Bool("help", false, "prints this message")
        login = flag.Bool("login", false, "use api_key to get & save access token")
        loginas = flag.String("loginas", "", "use api_key of client to get & save access token")
        ping = flag.Bool("ping", false, "use access token to ping hes")
        debug = flag.Bool("debug", false, "set to true for debug info")
        path = flag.String("path", "clients", "set to desired api path, eg. path=clients")
)

func main() {
        flag.Parse()
        if flag.NFlag() >=2 { // if number of flags >=2 then one of them must be -debug
                if flag.NFlag() == 2 {
                        if !*debug {
                                fmt.Println("Too many flags. Usage: ")
                                os.Exit(1)
                        }
                } else {
                        fmt.Println("Too many flags. Usage: ")
                        os.Exit(1)
                }
        }
        if *help {
                fmt.Println("Guru Clients ...")
                flag.PrintDefaults()
                os.Exit(1)
        }

        raw, err := hes.Asset("static/apiKey")
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
        raw = raw[0:len(raw)-1]
        at, err := hes.Asset("static/accessToken")
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
        //at = at[0:len(at)-1]

        kb,err := json.Marshal(hes.ApiKey{string(raw)})
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        if *ping {
                fmt.Println("Pinging hes ...")
                *path = "ping"
        } else if *login {
                fmt.Println("Logging in ...")
                *path = "tokens"
        } else if *loginas != "" {
                _, err := hes.Asset("static/apiKeys")
                if err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                }
        } else {

        }

        c := data.P.Get() //c, _ := data.RedisConn()
        defer c.Close()

        hesLive := []byte(apiHes)
        doneChan := make(chan bool, 1)
        errChan := make(chan error, 1)
        respChan := make(chan []byte, 1)
        fmt.Println("Getting ...")
        var url []byte
        switch *path {
        case "ping":
                url = append(hesLive, *path...)
                if *debug {
                        fmt.Println("url: ", url)
                }
        case "tokens":
                url = append(hesLive, *path...)
                if *debug {
                        fmt.Println("url: ", string(url))
                }
        case "clients":
                //url = append(hesLive, *path...)
                url = append(hesLive, "clients?page=2&limit=50"...)
                if *debug {
                        fmt.Println("url: ", string(url))
                }
        case "client":
                url = append(hesLive, "clients/2"...)
                if *debug {
                        fmt.Println("url: ", string(url))
                }
        default:
                fmt.Println("No request ...")
                os.Exit(1)
        }
        go func(url string) {
                req, err := http.NewRequest("GET", url, nil)
                if *login {
                        br := bytes.NewReader(kb)
                        req, err = http.NewRequest("POST", url, br)
                        hes.SetLoginHeaders(req)
                        test := make([]byte, 100)
                        bod,_ := req.GetBody()
                        _,_ = bod.Read(test)
                        fmt.Fprintf(os.Stderr,"%v\n", string(test))
                } else {
                        ato := &hes.AccessToken{}
                        if err = json.Unmarshal(at, ato); err != nil {
                                fmt.Println("Unmarshal access token error: ", err)
                                os.Exit(1)
                        }
                        hes.SetAccessHeaders(req, ato.Access_token) //hes.SetAccessHeaders(req, string(at))
                }
                if err != nil {
                        errChan <- err
                        return
                }
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
                                        var bb bytes.Buffer
                                        nn, err := bb.ReadFrom(resp.Body)
                                        if err != nil {
                                                //errChan <- err
                                                errChan <- errors.New("default "+err.Error())
                                                 return
                                        }
                                        fmt.Println("Bytes recd.: ", nn, resp.Status)
                                        respChan <- bb.Bytes()
                                         return

                        }
                        //respChan <- []byte(ct)
                        //return
                }
                doneChan<- true
        }(string(url))

        hcr := &hes.HesClientsResp{}
        select {
        case <-doneChan:
                fmt.Println("Done: ")
        case err = <-errChan:
                fmt.Println("Error: ", err)
        case bs := <-respChan:
                switch *path {
                case "tokens":
                        ioutil.WriteFile(hestkn, bs, os.ModePerm)
                        fmt.Println(string(bs))
                case "ping":
                        fmt.Println(string(bs))
                case "clients":
                        fmt.Println("Unmarshaling")
                        err := json.Unmarshal(bs, hcr)
                        if err != nil {
                                fmt.Println(err)
                                errChan <- err
                                return
                        }
                }
        case <-time.After(5000 * time.Millisecond):
                fmt.Println("Timeout: timed out")
        }

        close(doneChan)
        close(errChan)
        close(respChan)

        fmt.Println(hcr)
        hcs := hes.HesClients(hcr.Clients)
        sort.Sort(hcs)
        for _, hc := range hcs {
            fmt.Println("Clients...")
            fmt.Println(hc.String())
            //c.Do("SET", 
        }

}

