package hes

import (
        "fmt"
)

type HesClient struct {
	Client_id          int64 `json:"client_id,omitempty"`
	Name       string `json:"name,omitempty"`
	Rate_limit int64 `json:"rate_limit,omitempty"`
	Created_at     string `json:"created_at,omitempty"`
	Updated_at     string `json:"updated_at,omitempty"`
}

type HesLink struct {
	Self string `json:"self,omitempty"`
	Sites string `json:"sites,omitempty"`
	Gateways string `json:"gateways,omitempty"`
	Hubs string `json:"hubs,omitempty"`
	Accounts string `json:"accounts,omitempty"`
	Readings string `json:"readings,omitempty"`
	Transactions string `json:"transactions,omitempty"`
	Tariffs string `json:"tariffs,omitempty"`
}

func (c *HesClient) HeaderString() string{

	return fmt.Sprintf(" %10s %v %8s %10s %10s %8s %10s %10s %10s", "Client_id", "Name", "Rate_limit", "Created_at", "Updated_at")

}

func (c *HesClient) String() string{

	return fmt.Sprintf(" %10s %v %8s %10s %10s %8s %10s %10s %10s", c.Client_id, c.Name, c.Rate_limit, c.Created_at, c.Updated_at)

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

