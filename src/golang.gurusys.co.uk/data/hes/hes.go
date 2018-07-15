package hes

import (
        "fmt"
        "net/http"
)

type ApiKey struct {
	Api_key          string `json:"api_key,omitempty"`
}

type AccessToken struct {
	Access_token          string `json:"access_token,omitempty"`
}

type HesClientsResp struct {
        Links map[string]string `json:"_links",omitempty"`
        Clients []HesClient `json:"clients",omitempty"`
}

type HesClient struct {
	Client_id          int64 `json:"client_id,omitempty"`
	Name       string `json:"name,omitempty"`
	Rate_limit int64 `json:"rate_limit,omitempty"`
	Created_at     string `json:"created_at,omitempty"`
	Updated_at     string `json:"updated_at,omitempty"`
        Links map[string]map[string]string `json:"_links,omitempty"`
        //Links []HesLink `json:"_links,omitempty"`
}

type HesLink struct {
	Self []string `json:"self,omitempty"`
	Sites []string `json:"sites,omitempty"`
	Gateways []string `json:"gateways,omitempty"`
	Hubs []string `json:"hubs,omitempty"`
	Accounts []string `json:"accounts,omitempty"`
	Readings []string `json:"readings,omitempty"`
	Transactions []string `json:"transactions,omitempty"`
	Tariffs []string `json:"tariffs,omitempty"`
}

func (c *HesClient) HeaderString() string{
	return fmt.Sprintf(" %8s %8s %8s %20s %20s \n", "Client_id", "Name", "Rate_limit", "Created_at", "Updated_at")

}

func (c *HesClient) String() string{
	d := fmt.Sprintf(" %8d %8s %8d %20s %20s \n", c.Client_id, c.Name, c.Rate_limit, c.Created_at, c.Updated_at)
        var e string
        for k,v := range c.Links {
                e += fmt.Sprintf("%10s %v \n", k, v)
        }
	return fmt.Sprintf(" %s %s \n", d, e)

}

//HesClients implements sort.Interface
type HesClients []HesClient

func (cs HesClients) Len() int {
        return len(cs)
}

func (cs HesClients) Swap(i,j int) {
        cs[i], cs[j] = cs[j], cs[i]
}

func (cs HesClients) Less(i, j int) bool {
    return cs[i].Name < cs[j].Name
}

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
