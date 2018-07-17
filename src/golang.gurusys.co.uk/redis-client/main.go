package main

import (
        "bufio"
        "bytes"
        "compress/gzip"
        "errors"
        "flag"
        "fmt"
        _"encoding/gob"
        "net/http"
        _"io/ioutil"
        "encoding/json"
        "os"
        _"os/signal"
        _"sort"
        "strconv"
        "strings"
        _"sync"
        "time"

        "golang.gurusys.co.uk/data/pw"
        _"golang.gurusys.co.uk/data"
        "golang.gurusys.co.uk/getter"
)

var (
        err error
        help = flag.Bool("help", false, "prints this message")
        ping = flag.Bool("ping", false, "use access token to ping hes")
        debug = flag.Bool("debug", false, "set to true for debug info")
        path = flag.String("path", "account", "set to desired api path, eg. path=account")
)


func main() {
        intc := make(chan os.Signal, 1)
        getter.SignalNotify(intc)

        flag.Parse()

        cmds := make(chan string, 1)
        resc := make(chan []byte, 1)
        errc := make(chan error, 1)
        page := make(map[string]int32)
        page["page_number"] = 1
        page["page_size"] = 10
        pb, err := json.Marshal(page)
        if err != nil {
                fmt.Println("Json page: ", err)
                os.Exit(1)
        }
        go func() {
        execute:
                a:=<-cmds
                a = strings.TrimSpace(a)
                if strings.Contains(a, " ") {
                        args := strings.Split(a, " ")
                        if len(args) != 2 {
                                goto execute
                        }
                        switch args[0]{
                        case "user","lead","company","person","opportunity","project","task","activity":
                                fmt.Printf("executing %s\n", url+args[0]+"s/"+args[1] )
                                req, err := http.NewRequest("GET", url+args[0]+"s/"+args[1], nil)
                                pw.SetAccessHeaders(req, keys["pw"].Api_key)
                                resp, err := getter.Client.Do(req)
                                if err != nil {
                                        errc <- errors.New("GET"+err.Error())
                                        goto execute
                                }else {
                                        if resp != nil {
                                                defer resp.Body.Close()
                                        }
                                        cl := resp.Header.Get(getter.ContentLengthHeader)
                                        icl := 1000
                                        if cl != "" {
                                                icl, err = strconv.Atoi(cl)
                                                if err != nil {
                                                        //errc <- err
                                                        errc <- errors.New("Strconv "+err.Error())
                                                        icl = 1000//return
                                                }
                                        }
                                        ubs := make([]byte, icl*3)
                                        ct := resp.Header.Get(getter.ContentTypeHeader)
                                        switch ct {
                                                default: //case "gzip": // Prosperworks doesn't inform that content is gzipped
                                                        gzr, err := gzip.NewReader(resp.Body)
                                                        if err != nil {
                                                                 //errc <- err
                                                                errc <- errors.New("gzip"+err.Error())
                                                                goto execute //return
                                                        }
                                                        defer gzr.Close()
                                                        nbs, err := gzr.Read(ubs)
                                                        ubs = ubs[:nbs]
                                                        fmt.Println("Bytes recd.: ", nbs, resp.Status)
                                                        resc <- ubs
                                                /*default:
                                                        var bb bytes.Buffer
                                                        nn, err := bb.ReadFrom(resp.Body)
                                                        if err != nil {
                                                                //errc <- err
                                                                errc <- errors.New("default "+err.Error())
                                                                 return
                                                        }
                                                        fmt.Println("Bytes recd.: ", nn, resp.Status)
                                                        resc <- bb.Bytes()
                                                */
                                        }
                                        goto execute // return
                                        //respChan <- []byte(ct)
                                        //return
                                }
                                //resc<-[]byte("done "+a)
                        default:
                                resc<-[]byte("unknown "+a)
                                goto execute // return
                        }
                } else {
                        switch a{
                        case "exit":
                                fmt.Printf("exiting %s\n", a)
                                resc<-[]byte("done")
                                return
                        case "account":
                                fmt.Printf("executing %s\n", a)
                                req, err := http.NewRequest("GET", url+a, nil)
                                pw.SetAccessHeaders(req, keys["pw"].Api_key)
                                resp, err := getter.Client.Do(req)
                                if err != nil {
                                        errc <- errors.New("GET"+err.Error())
                                        goto execute
                                }else {
                                        if resp != nil {
                                                defer resp.Body.Close()
                                        }
                                        cl := resp.Header.Get(getter.ContentLengthHeader)
                                        icl := 100
                                        if cl != "" {
                                                icl, err = strconv.Atoi(cl)
                                                if err != nil {
                                                        //errc <- err
                                                        errc <- errors.New("Strconv "+err.Error())
                                                        icl = 100//return
                                                }
                                        }
                                        ubs := make([]byte, icl*3)
                                        ct := resp.Header.Get(getter.ContentTypeHeader)
                                        switch ct {
                                                default: //case "gzip": // Prosperworks doesn't inform that content is gzipped
                                                        gzr, err := gzip.NewReader(resp.Body)
                                                        if err != nil {
                                                                 //errc <- err
                                                                errc <- errors.New("gzip"+err.Error())
                                                                goto execute
                                                        }
                                                        defer gzr.Close()
                                                        nbs, err := gzr.Read(ubs)
                                                        ubs = ubs[:nbs]
                                                        resc <- ubs
                                                /*default:
                                                        var bb bytes.Buffer
                                                        nn, err := bb.ReadFrom(resp.Body)
                                                        if err != nil {
                                                                //errc <- err
                                                                errc <- errors.New("default "+err.Error())
                                                                 return
                                                        }
                                                        fmt.Println("Bytes recd.: ", nn, resp.Status)
                                                        resc <- bb.Bytes()
                                                */
                                        }
                                        goto execute // return
                                        //respChan <- []byte(ct)
                                        //return
                                }
                                //resc<-[]byte("done "+a)
                        case "users","leads","companies","people","opportunities","projects","tasks","activities":
                                fmt.Printf("executing %s\n", a)
                                br := bytes.NewReader(pb)
                                req, err := http.NewRequest("POST", url+a+search, br)
                                pw.SetAccessHeaders(req, keys["pw"].Api_key)
                                resp, err := getter.Client.Do(req)
                                if err != nil {
                                        errc <- errors.New("GET"+err.Error())
                                        goto execute
                                }else {
                                        if resp != nil {
                                                defer resp.Body.Close()
                                        }
                                        cl := resp.Header.Get(getter.PwResponseHeader)
                                        icl := 5000
                                        if cl != "" {
                                                icl, err = strconv.Atoi(cl)
                                                if err != nil {
                                                        //errc <- err
                                                        errc <- errors.New("Strconv "+err.Error())
                                                        icl = 100//return
                                                }
                                        }
                                        ubs := make([]byte, 15000)//icl*3)
                                        ct := resp.Header.Get(getter.ContentTypeHeader)
                                        switch ct {
                                                default: //case "gzip": // Prosperworks doesn't inform that content is gzipped
                                                        gzr, err := gzip.NewReader(resp.Body)
                                                        if err != nil {
                                                                 //errc <- err
                                                                errc <- errors.New("gzip"+err.Error())
                                                                goto execute //return
                                                        }
                                                        defer gzr.Close()
                                                        nbs, err := gzr.Read(ubs)
                                                        ubs = ubs[:nbs]
                                                        fmt.Println("Bytes recd.: ", nbs, resp.Status, icl)
                                                        resc <- ubs
                                                /*default:
                                                        var bb bytes.Buffer
                                                        nn, err := bb.ReadFrom(resp.Body)
                                                        if err != nil {
                                                                //errc <- err
                                                                errc <- errors.New("default "+err.Error())
                                                                 return
                                                        }
                                                        fmt.Println("Bytes recd.: ", nn, resp.Status)
                                                        resc <- bb.Bytes()
                                                */
                                        }
                                        goto execute // return
                                        //respChan <- []byte(ct)
                                        //return
                                }
                                //resc<-[]byte("done "+a)
                        default:
                                resc<-[]byte("unknown "+a)
                                goto execute // return
                        }
                }
        }()
read:
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Cmd: ")
        cmd, _ := reader.ReadString('\n')
        if len(cmd)<2 {
                fmt.Println("exit|account|users|lead|companies|people|projects|tasks|activities")
                fmt.Println("--OR--")
                fmt.Println("user <id>|lead <id>|company <id>|people <id>|project <id>|task <id>")
                goto read
        }
        cmds<-cmd
        cmd = strings.TrimSuffix(cmd, "\n")

        select {
        case r:=<-resc:
                if string(r)=="done" {
                        goto end //return
                }
                switch cmd {
                case "account":
                        fmt.Printf("Resc: %s\n", string(r))
                case "users":
                        fmt.Printf("Users: %d\n", 20)
                        var users []pw.User
                        if err = json.Unmarshal(r, &users); err != nil {
                                errc<- err
                        }
                        for _, u := range users {
                                fmt.Println(u)
                        }
                        goto read
                case "leads":
                        var leads pw.Leads
                        if err = json.Unmarshal(r, &leads); err != nil {
                                errc<- err
                        }
                        fmt.Print(leads.String())
                        for _, l := range leads {
                                fmt.Println(l.String())
                        }
                        goto read
                case "companies":
                        var cs pw.Companies
                        if err = json.Unmarshal(r, &cs); err != nil {
                                errc<- err
                        }
                        fmt.Print(cs.String())
                        for _, c := range cs {
                                fmt.Println(c.String())
                        }
                        goto read
                case "people":
                        var ps pw.People
                        if err = json.Unmarshal(r, &ps); err != nil {
                                errc<- err
                        }
                        fmt.Print(ps.String())
                        for _, p := range ps {
                                fmt.Println(p.String())
                        }
                        goto read
                case "opportunities":
                        var os pw.Opportunities
                        if err = json.Unmarshal(r, &os); err != nil {
                                errc<- err
                        }
                        fmt.Print(os.String())
                        for _, o := range os {
                                fmt.Println(o.String())
                        }
                        goto read
                case "projects":
                        var ps pw.Projects
                        if err = json.Unmarshal(r, &ps); err != nil {
                                errc<- err
                        }
                        fmt.Print(ps.String())
                        for _, p := range ps {
                                fmt.Println(p.String())
                        }
                        goto read
                case "tasks":
                        var ts pw.Tasks
                        if err = json.Unmarshal(r, &ts); err != nil {
                                errc<- err
                        }
                        fmt.Print(ts.String())
                        for _, t := range ts {
                                fmt.Println(t.String())
                        }
                        goto read
                case "activities":
                        var as pw.Activities
                        if err = json.Unmarshal(r, &as); err != nil {
                                errc<- err
                        }
                        fmt.Print(as.String())
                        for _, a := range as {
                                fmt.Println(a.String())
                        }
                        goto read
                default:
                        fmt.Println(string(r))
                }
        case err=<-errc:
                fmt.Println("Errc: ", err)
                goto read
        case <-intc:
                cmds<-"exit"
                goto end
        case <-time.After(30 * time.Second):
                cmds<-"exit"
                fmt.Println("Timeout: timed out")
                goto end
        }
        goto read

end:
        close(cmds)
        close(resc)
        close(intc)

}
