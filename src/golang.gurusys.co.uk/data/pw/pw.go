package pw

import (
        _"fmt"
        "net/http"
)

func SetLoginHeaders(req *http.Request) {

	//req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; rv:39.0) Gecko/20100101 Firefox/39.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("DNT", "1")
	req.Header.Set("Host", "hes.gurusys.co,uk")
	req.Header.Set("Referer", "hes.gurusys.co.uk")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:31.0) Gecko/20100101 Firefox/59.0")
        req.Header.Set("X-Requested-With", "XMLHttpRequest")
	//req.Header.Set("Cache-Control", "max-age=0")
	//this.Req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*,q=0.8")
}

func SetAccessHeaders(req *http.Request, accesstoken string) {
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; rv:39.0) Gecko/20100101 Firefox/39.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Authorization", "Bearer " + accesstoken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("DNT", "1")
	req.Header.Set("Host", "hes.gurusys.co,uk")
	req.Header.Set("Referer", "hes.gurusys.co.uk")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:31.0) Gecko/20100101 Firefox/59.0")
        req.Header.Set("X-Requested-With", "XMLHttpRequest")
	//req.Header.Set("Cache-Control", "max-age=0")
	//this.Req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*,q=0.8")
}
