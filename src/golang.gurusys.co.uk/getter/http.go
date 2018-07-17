package getter

import (
	_"compress/gzip"
	"crypto/tls"
	_"encoding/json"
	_"errors"
	"fmt"
	"io"
	_ "log"
	"net"
	"net/http"
	stdurl "net/url"
	_ "os"
	_"reflect"
	_"strconv"
	"strings"
	"sync"
	_ "time"
)

var (
	tr                  = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	Client              = &http.Client{Transport: tr}
	ContentLengthHeader = "Content-Length"
	ContentTypeHeader   = "Content-Type"
	acceptRangesHeader  = "Accept-Ranges"
	status              = "Status"
        PwResponseHeader    = "X-Pw-Total"
	mu                  = &sync.Mutex{}
)

type HttpGetter struct {
	RespReader  io.ReadCloser
	Resp        *http.Response
	Req         *http.Request
	Clnt        *http.Client
	RespHeaders map[string]string
	Ubs         []byte
}

type HttpMultiGetter struct {
	sync.RWMutex
	RespReader  []io.ReadCloser
	Resp        []*http.Response
	Req         []*http.Request
	Clnt        *http.Client
	RespHeaders []map[string]string
	Ubs         [][]byte
}

func NewHttpGetter(url string) (*HttpGetter, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return &HttpGetter{Req: req, Clnt: Client}, nil

}

func NewHttpMultiGetter(urls []string) (*HttpMultiGetter, error) {

	hmg := &HttpMultiGetter{Clnt: Client}
	hmg.Req = make([]*http.Request, len(urls))
	hmg.Resp = make([]*http.Response, len(urls))
	hmg.Ubs = make([][]byte, len(urls))
	hmg.RespReader = make([]io.ReadCloser, len(urls))
	for i, url := range urls {
		r, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		hmg.Req[i] = r
	}
	return hmg, nil

}

func (this *HttpGetter) Display() string {

	//switch t.(type) {
	//      case OptionData:
	return fmt.Sprintf(" %v \n", string(this.Ubs))
	//    default:
	//               return fmt.Sprintf("%v", this.Object)
	// }
	//fmt.Println(i, ". ", oD.Book[0].ExpiryDate, "\t", oD.Book[0].BuyQuantity1, "\t", oD.Book[0].BuyPrice1, "\t", oD.Book[0].SellPrice1, "\t", oD.Book[0].SellQuantity1, "\t", oD.Book[0].LowPrice, "\t", oD.Book[0].HighPrice)

}

func (this *HttpMultiGetter) Display() string {
	return fmt.Sprintf(" %v \n", this.Ubs)
}

/*func (this *HttpGetter) Unmarshal(v interface{}) error {
	err := json.Unmarshal(this.Ubs, v)
	if err != nil {
		return err //errors.New("Json unmarshalling error")
	}
	switch v.(type) {
	case *data.OptionData:
		this.Od = v.(*nse.OptionData)
	}
	return nil
}

func (this *HttpMultiGetter) MultiUnmarshal(v interface{}) error {

	for i, ubs := range this.Ubs {
		//err := json.Unmarshal(ubs, v)
		//if err != nil {
		//        return err //errors.New("Json unmarshalling error")
		// }
		switch v.(type) {
		case *nse.OptionData:
			if i > len(this.Ods)-1 {
				return errors.New(fmt.Sprintf("Unmarshal error: %v", i))
			}
			this.Ods[i] = &nse.OptionData{}
			err := json.Unmarshal(ubs, this.Ods[i])
			if err != nil {
				return err //errors.New("Json unmarshalling error")
			}
			//*this.Ods[i] = *v.(*OptionData)
		default:
			return errors.New(fmt.Sprintf("Unrecognized type %v", reflect.TypeOf(v)))
		}
	}
	return nil

}

func (this *HttpGetter) Get() error {

	this.SetHeaders()
	x, err := this.Clnt.Do(this.Req)
	this.Resp = x
	defer this.Resp.Body.Close()
	if err != nil {
		return err
	}
	cl := this.Resp.Header.Get(ContentLengthHeader)
	if cl == "" {
		return errors.New("Response doesn't have content length")
	}
	icl, err := strconv.Atoi(cl)
	if err != nil {
		return err
	}
	this.Ubs = make([]byte, icl*3)
	ct := this.Resp.Header.Get(ContentTypeHeader)
	if ct == "" {
		return errors.New("Response doesn't have content type")
	}
	switch ct {
	case "gzip":
		this.RespReader, err = gzip.NewReader(this.Resp.Body)
		defer this.RespReader.Close()
		if err != nil {
			return err
		}
	default:
		this.RespReader, err = gzip.NewReader(this.Resp.Body)
		defer this.RespReader.Close()
		if err != nil {
			return err
		}
	}
	n, err := this.RespReader.Read(this.Ubs)
	if err != nil && err != io.EOF {
		return err
	}
	this.Ubs = this.Ubs[:n]
	return nil

}

func (this *HttpGetter) MultiGet(doneChan chan bool, errorChan chan error) error {
	go func(d *HttpGetter) {
		this.SetHeaders()
		x, err := this.Clnt.Do(this.Req)
		this.Resp = x
		defer this.Resp.Body.Close()
		if err != nil {
			errorChan <- err
		}
		cl := this.Resp.Header.Get(ContentLengthHeader)
		if cl == "" {
			errorChan <- errors.New("Response doesn't have content length")
		}
		icl, err := strconv.Atoi(cl)
		if err != nil {
			errorChan <- err
		}
		this.Ubs = make([]byte, icl*3)
		ct := this.Resp.Header.Get(ContentTypeHeader)
		if ct == "" {
			errorChan <- errors.New("Response doesn't have content type")
		}
		switch ct {
		case "gzip":
			this.RespReader, err = gzip.NewReader(this.Resp.Body)
			defer this.RespReader.Close()
			if err != nil {
				errorChan <- err
			}
		default:
			this.RespReader, err = gzip.NewReader(this.Resp.Body)
			defer this.RespReader.Close()
			if err != nil {
				errorChan <- err
			}
		}
		n, err := this.RespReader.Read(this.Ubs)
		if err != nil {
			errorChan <- err
		}
		this.Ubs = this.Ubs[:n]
		doneChan <- true
	}(this)
	return nil
}

func (this *HttpMultiGetter) MultiGet(doneChan chan bool, errorChan chan error) error {
	var ws sync.WaitGroup
	for i, req := range this.Req {
		ws.Add(1)
		go func(d *HttpMultiGetter, i int, rqst *http.Request) {
			defer ws.Done()
			if i > len(d.Req) {
				ws.Done()
			}
			d.SetMultiHeaders(i)
			x, err := d.Clnt.Do(rqst) //can't use req?
			if err != nil {
				errorChan <- err
			}
			d.Resp[i] = x
			defer d.Resp[i].Body.Close()
			cl := d.Resp[i].Header.Get(ContentLengthHeader)
			if cl == "" {
				errorChan <- errors.New("Response doesn't have content length")
			}
			icl, err := strconv.Atoi(cl)
			if err != nil {
				errorChan <- err
			}
			d.Ubs[i] = make([]byte, icl*3)
			ct := d.Resp[i].Header.Get(ContentTypeHeader)
			if ct == "" {
				errorChan <- errors.New("Response doesn't have content type")
			}
			switch ct {
			case "gzip":
				d.RespReader[i], err = gzip.NewReader(d.Resp[i].Body)
				defer d.RespReader[i].Close()
				if err != nil {
					errorChan <- err
				}
			default:
				d.RespReader[i], err = gzip.NewReader(d.Resp[i].Body)
				defer d.RespReader[i].Close()
				if err != nil {
					errorChan <- err
				}
			}
			n, err := d.RespReader[i].Read(d.Ubs[i])
			if err != nil {
				errorChan <- err
			}
			d.Ubs[i] = d.Ubs[i][:n]
			//doneChan<- true
		}(this, i, req)
	}
	ws.Wait()
	doneChan <- true
	return nil
}

func (this *HttpGetter) SetHeaders() {
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; rv:39.0) Gecko/20100101 Firefox/39.0")
	this.Req.Header.Set("Accept", "* /*")
	this.Req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	this.Req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	this.Req.Header.Set("Connection", "keep-alive")
	this.Req.Header.Set("DNT", "1")
	this.Req.Header.Set("Host", "www.nseindia.com")
	this.Req.Header.Set("Referer", "https://www.nseindia.com")
	this.Req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:31.0) Gecko/20100101 Firefox/59.0")
    this.Req.Header.Set("X-Requested-With", "XMLHttpRequest")
	//req.Header.Set("Cache-Control", "max-age=0")
	//this.Req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,* /*,q=0.8")
}*/

/*func (this *HttpMultiGetter) SetMultiHeaders(i int) {
	this.Lock()
	defer this.Unlock()
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; rv:39.0) Gecko/20100101 Firefox/39.0")
	this.Req[i].Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:31.0) Gecko/20100101 Firefox/31.0 Iceweasel/31.8.0i")
	this.Req[i].Header.Set("Host", "nseindia.com")
	this.Req[i].Header.Set("DNT", "1")
	this.Req[i].Header.Set("Connection", "keep-alive")
	//req.Header.Set("Cache-Control", "max-age=0")
	this.Req[i].Header.Set("Accept-Language", "en-US,en;q=0.5")
	this.Req[i].Header.Set("Accept-Encoding", "gzip, deflate")
	this.Req[i].Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,* /*,q=0.8")
}*/

type HttpDownloader struct {
	url  string
	file string
	par  int64
	len  int64
	ips  []string
	//parts []Part
	skipTLS   bool
	resumable bool
}

func NewHttpDownloader(url string, par int64, skipTLS bool) *HttpDownloader {
	parsed, err := stdurl.Parse(url)
	FatalCheck(err)

	ips, err := net.LookupIP(parsed.Host)
	FatalCheck(err)

	ipstr := FilterIPV4(ips)
	fmt.Sprintf("Resolved host: %v \n", strings.Join(ipstr, "|"))

	req, err := http.NewRequest("GET", url, nil)
	FatalCheck(err)

	resp, err := Client.Do(req)
	FatalCheck(err)
	fmt.Println("Done")

	if resp.Header.Get(ContentLengthHeader) == "" {
		fmt.Println("Content Length not set")
	} else {
		fmt.Sprintf("Content Length: %q \n", resp.Header.Get(ContentLengthHeader))
	}

	if resp.Header.Get(acceptRangesHeader) == "" {
		fmt.Println("Accept ranges not set")
	}

	ret := new(HttpDownloader)
	return ret
}

func (d *HttpDownloader) Do(doneChan chan bool) {
	//time.Sleep(1000)
	doneChan <- true
}
