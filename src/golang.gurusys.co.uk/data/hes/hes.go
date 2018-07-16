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

type HesSitesResp struct {
        Links map[string]string `json:"_links",omitempty"`
        Sites []HesSite `json:"sites",omitempty"`
}

type HesSite struct {
        Client_id int64 `json:"client_id,omitempty"`// The unique ID of the client that this site belongs to
        Site_id int64 `json:"site_id,omitempty"` // The unique ID of the site itself
        Name string `json:"name,omitempty"` //  The name of this site
        Sbri_export_allowed bool `json:"sbri_export_allowed,omitempty"` // Whether or not this site will export to the SBRI system
        Ping_automatically bool `json:"ping_automatically,omitempty"` // Whether or not this site will be pinged automatically at a defined frequency
        Ping_frequency int32 `json:"ping_frequency,omitempty"` // If the site is configured to ping automatically, how many seconds there will be between pings
        Last_ping_at string `json:"last_ping_at,omitempty"` //The date on which the last ping occurred, in ISO8601 format
        Gateway_count int32 `json:"gateway_count,omitempty"` // How many gateways are associated with this site
        Hub_count int32 `json:"hub_count,omitempty"` // How many hubs are associated with this site
        Created_at string `json:"created_at,omitempty"` // When this site record was created, in ISO8601 format
        Updated_at string  `json:"updated_at,omitempty"` // When this site record was last changed, in ISO8601 format
        Links map[string]map[string]string `json:"_links,omitempty"`
}

func (s *HesSite) HeaderString() string{
	return fmt.Sprintf(" %8d %8d %8s %v %2d %d \n", "Client_id", "Site_Id", "Name", "Sbri?", "Gateways", "Hubs")

}

func (s *HesSite) String() string{
	d := fmt.Sprintf(" %8d %8d 8s %v %d %d \n", s.Client_id, s.Site_id, s.Name, s.Sbri_export_allowed, s.Gateway_count, s.Hub_count)
        var e string
        for k,v := range s.Links {
                e += fmt.Sprintf("%10s %v \n", k, v)
        }
	return fmt.Sprintf(" %s %s \n", d, e)

}

type HesGatewaysResp struct {
        Links map[string]string `json:"_links",omitempty"`
        Gateways []HesGateway `json:"gateways",omitempty"`
}

type HesGateway struct {
        Client_id int64 `json:"client_id,omitempty"` //  The unique ID of the client that this gateway belongs to
        Site_id int64 `json:"site_id,omitempty"` // The unique ID of the site that this gateway belongs to
        Gateway_id int64 `json:"gateway_id,omitempty"` // The unique ID of the gateway itself
        Name string `json:"name,omitempty"` // The name of this gateway
        Ip_address string `json:"ip_address,omitempty"`//  The IP address of this gateway
        Ip_port int32 `json:"ip_port"` // The port over which this gateway is communicating
        Local_mac_address string `json:"local_mac_address"` // The MAC address of this gateway
        Last_network_crawl_completed_at string `json:"last_network_crawl_completed_at,omitempty"` // The date on which the last network crawl was completed for this gateway, in ISO8601 format
        Network_crawl_in_progress bool `json:"network_crawl_in_progress,omitempty"` // Whether or not a network crawl is currently in progress
        Node_count int64 `json:"node_count,omitempty"`// How many nodes are on this gateway’s network
        Software_build_time string `json:"software_build_time,omitempty"`// When the software for this gateway was created, in ISO8601 format
        Contactable bool `json:"contactable,omitempty"` // Whether this gateway is currently contactable; reserved for later use
        Network_type int32 `json:"network_type,omitempty"`// Which type of network this gateway is connected to; currently 2400 for a 2.4GHz network
        Sbri_export_allowed bool `json:"sbri_export_allowed"` // Whether this gateway’s data is allowed to be used in the SBRI system
        Created_at string  `json:"created_at,omitempty"` // When this gateway record was created, in ISO8601 format
        Updated_at string `json:"updated_at,omitempty"`// When this gateway record was last changed, in ISO8601 format
        Links map[string]map[string]string `json:"_links,omitempty"`
}

func (g *HesGateway) HeaderString() string{
	return fmt.Sprintf(" %8s %8s %8s %8s %8s %8s %8s \n", "Client_id", "Site_Id", "Gateway_Id", "Name", "Sbri?", "Ip_Address", "Ip_port")
}

func (g *HesGateway) String() string{
	d := fmt.Sprintf(" %8d %8d %8d 8s %v %d %d \n", g.Client_id, g.Site_id, g.Gateway_id, g.Name, g.Sbri_export_allowed, g.Ip_address, g.Ip_port)
        var e string
        for k,v := range g.Links {
                e += fmt.Sprintf("%10s %v \n", k, v)
        }
	return fmt.Sprintf(" %s %s \n", d, e)
}

type HesHubResp struct {
        Links map[string]string `json:"_links",omitempty"`
        Hubs []HesHub `json:"hubs",omitempty"`
}

type HesHub struct {
        Client_id int64 `json:"client_id,omitempty"`// The unique ID of the client that this hub belongs to
        Site_id int64 `json:"site_id,omitempty"`// The unique ID of the site that this hub belongs to
        Gateway_id int64 `json:"gateway_id,omitempty"`// The unique ID of the gateway that this hub belongs to
        Hub_id string `json:"hub_id,omitempty"` //  The unique ID of the hub itself
        Relays []Relay `json:"relays,omitempty"` // An array of data about the relays associated with this hub, contents detailed below
        Location  []Location `json:"location,omitempty"` // An array of data about the location associated with this hub, contents detailed below
        Analog_input_reporting_interval int32 `json:"analog_input_reporting_interval,omitempty"`// How frequently this hub is programmed to send analog input reports
        Awaiting_configuration bool `json:"awaiting_configuration,omitempty"`// Whether this hub is currently awaiting configuration
        Last_serial_user string `json:"last_serial_user,omitempty"`// The name of the last user who connected to this hub via the serial port
        Zigbee_allow_joining bool `json:"zigbee_allow_joining,omitempty"` // Whether or not Zigbee devices are allowed to connect to this hub
        Mbus_baudrate string `json:"mbus_baudrate,omitempty"`// The baud rate of the Mbus network used by this hub
        Contactable bool `json:"is_contactable,omitempty"` // Whether or not this hub is currently accessible over the network. Reserved for future development.
        Created_at string  `json:"created_at,omitempty"` // When this hub was first seen by the HES, in ISO8601 format
        Updated_at string `json:"updated_at,omitempty"`//When this hub was last updated by the HES, in ISO8601 format
        Links map[string]map[string]string `json:"_links,omitempty"`
}

func (h *HesHub) HeaderString() string{
	return fmt.Sprintf(" %8s %8s %8s %8s %8s %8s %8s \n", "Client_id", "Site_Id", "Gateway_Id", "Hub_Id", "Relays", "Location", "Last_serial_user")
}

func (h *HesHub) String() string{
	d := fmt.Sprintf(" %8d %8d %8d %8s %d %s %s \n", h.Client_id, h.Site_id, h.Gateway_id, h.Hub_id, len(h.Relays), h.Location[0].Name, h.Last_serial_user)
        var e string
        for k,v := range h.Links {
                e += fmt.Sprintf("%10s %v \n", k, v)
        }
	return fmt.Sprintf(" %s %s \n", d, e)
}

type Relay struct {
        Source string `json:"source,omitempty"` // info source
        State bool `json:"state,omitempty"` // active ?
}

type Location struct {
        Id int64 `json:"id,omitempty"`
        Name string `json:"name,omitempty"` // name
}

type HesHubAccountResp struct {
        Links map[string]string `json:"_links",omitempty"`
        Accounts []HesHubAccount `json:"accounts",omitempty"`
}

type HesHubAccount struct {
        Client_id int64 `json:"client_id,omitempty"`// The unique ID of the client that this account belongs to
        Site_id int64 `json:"site_id,omitempty"`// The unique ID of the site that this account belongs to
        Gateway_id int64 `json:"gateway_id,omitempty"`// The unique ID of the gateway that this account belongs to
        Hub_id string `json:"hub_id,omitempty"` //  The unique ID of the hub that this account belongs to
        Account_id int64 `json:"account_id,omitempty"`// The unique ID of the account that this account belongs to
        Account_reference string `json:"account_reference,omitempty"`// The human-readable account number for this account
        Credit_statuses []string `json:"credit_statuses,omitempty"`// An array of strings identifying what the credit status is for this account
        Last_read float32 `json:"last_reading_summation_kwh,omitempty"`// What the last reading was for this account in kWh
        Last_read_fmt float32 `json:"last_reading_summation_formatted,omitempty"`//What the last reading was for this account
        Last_read_units float32 `json:"last_reading_summation_unit_of_measure,omitempty"`//The unit of measure for the last reading
        Last_read_timestamp string `json:"last_reading_timestamp,omitempty"`//The timestamp of the last reading, in ISO8601 format
        Last_read_balance_gbp float32 `json:"last_reading_balance_gbp,omitempty"`//`The credit balance as of the last reading, in GBP
        Meter_serial_number string `json:"meter_serial_number,omitempty"`// The serial number of the meter associated with this account
        Payment_control_method string `json:"payment_control_method,omitempty"` //What payment control method is currently set for this account
        Pending_transaction_count int32 `json:"pending_transaction_count,omitempty"` //How many pending transactions there are for this account (those where confirmation of the transaction has not yet been sent from the hub)
        Pending_transaction_total float32 `json:"pending_transaction_total,omitempty"` //The total value of any pending transactions
        Pending_transactions []HesHubPendingTransaction `json:"pending_transactions,omitempty"` //An array of details about any pending transactions, documented below
        Unit_of_measure string `json:"unit_of_measure,omitempty"` //A text description of the unit of measure used for this account
        Temperature_unit_of_measure string `json:"temperature_unit_of_measure,omitempty"` //A text description of the unit of measure used for temperatures on this account
        Energy_carrier_unit_of_measure string `json:"energy_carrier_unit_of_measure,omitempty"` //A text description of the unit of measure used by the energy carrier for this account
        Meter_type string `json:"meter_type,omitempty"` //What type of energy is being measured for this account
        Meter_status string  `json:"meter_status,omitempty"`//The status of the attached meter
        Minimum_read_interval int32 `json:"minimum_read_interval,omitempty"` //How often meter readings must be received for this account, in seconds
        Current_debt_amount float32 `json:"current_debt_amount,omitempty"` //How much debt is set for this account
        Percentage_taken_as_debt_repayment float32 `json:"percentage_taken_as_debt_repayment,omitempty"` //What percentage of top-ups are applied against the debt balance of this account, if any
        Recovery_start_date string `json:"recovery_start_date,omitempty"` //When debts started to be collected for this account
        Emergency_credit_allowance float32 `json:"emergency_credit_allowance,omitempty"`  //What the account’s emergency credit allowance is
        Is_active bool `json:"is_active,omitempty"`       // Whether or not this account is currently active
        Active_from string `json:"active_from,omitempty"` //When this account became active, in ISO8601 format
        Expires_at string `json:"expires_at,omitempty"` //When this account will no longer be active, in ISO8601 format. If this is null, then the account is not currently set to expire
        Source []Src `json:"source,omitempty"`//An array of further details about this account, documented below
        Reporting_interval int32 `json:"reporting_interval,omitempty"` // How often this account is supposed to report data to the HES, in seconds
        Flow_temp float32 `json:"flow_temperature,omitempty"`// What the current inlet temperature was at last report
        Flow_temp_fmt float32 `json:"flow_temperature_formatted,omitempty"`// What the current inlet temperature was at last report, in the correct units
        Float_temp_units string `json:"flow_temperature_unit_of_measure,omitempty"`//The unit of measure for the current inlet temperature
        Return_temp float32 `json:"return_temperature,omitempty"` // What the current outlet temperature was at last report
        Return_temp_fmt float32 `json:"return_temperature_formatted,omitempty"`// What the current outlet temperature was at last report, in the correct units
        Return_temp_units string `json:"return_temperature_unit_of_measure,omitempty"`//The unit of measure for the current outlet temperature
        Flow_rate float32 `json:"flow_rate,omitempty"` //What the current flow rate was at last report
        flow_rate_formatted float32 `json"flow_rate_formatted"` //What the current flow rate was at last report, in the correct units
        Flow_rate_units string `json:"flow_rate_unit_of_measure,omitempty"`//The unit of measure for the current flow rate
        Created_at string  `json:"created_at,omitempty"` // When this account was first seen by the HES, in ISO8601 format
        Updated_at string `json:"updated_at,omitempty"`//When this account was last updated by the HES, in ISO8601 format
        Links map[string]map[string]string `json:"_links,omitempty"`
}

type HesHubPendingTransaction struct {
        Client_id int64 `json:"client_id,omitempty"`// The unique ID of the client that this transaction belongs to
        Site_id int64 `json:"site_id,omitempty"`// The unique ID of the site that this transaction belongs to
        Gateway_id int64 `json:"gateway_id,omitempty"`// The unique ID of the gateway that this transaction belongs to
        Hub_id string `json:"hub_id,omitempty"` //  The unique ID of the hub that this transaction belongs to
        Account_id int64 `json:"account_id,omitempty"`// The unique ID of the account that this transaction belongs to
        Account_reference string `json:"account_reference,omitempty"`// The human-readable account number for the account this transaction belongs to
        Transaction_id  int64  `json:"transaction_id,omitempty"` // The unique ID of the transaction
        Transaction_sequence int64 `json:"transaction_sequence,omitempty"`// The identifier for this transaction on the hub
        Amount string `json:"amount,omitempty"`//The amount of the transaction
        Type string `json:"type,omitempty"` //The type of transaction; this will be one of Credit: Set Balance, Credit: Incremental Adjustment, Credit: Top-Up, Debt: Set Debt, Debt: Incremental Adjustment
        Source string `json:"source,omitempty"` //Where this transaction was initiated from; this will be one of HES, Hub, API.
        Topup_token string `json:"topup_token,omitempty"` //For top-ups, this is the generated top-up token which can be entered on the hub.
        Client_reference string `json:"client_reference,omitempty"` //For transactions initiated via the API, the unique ID supplied by the client system; this is a string of maximum 31 characters in length
        Received_date string `json:"received_date,omitempty"` //When this transaction was received by the HES, in ISO8601 format
        Applied_date string `json:"applied_date,omitempty"` //When this transaction was successfully applied by the hub, in ISO8601 format. If the transaction has not yet been successfully applied, this could be null
        Links map[string]map[string]string `json:"_links,omitempty"`
}

type Src struct {
        Hub_id string `json:"hub_id,omitempty"` //The unique ID of the hub this account is associated with
        Location string `json:"location,omitempty"` //The location name associated with this account
        Scheme_name string `json:"scheme_name,omitempty"` //The name of the scheme associated with this account
        Scheme_id int64 `json:"scheme_id,omitempty"` //The ID of the scheme associated with this associated
}

type HesHubReading struct {
        Client_id int64 `json:"client_id,omitempty"`// The unique ID of the client that this reading belongs to
        Site_id int64 `json:"site_id,omitempty"`// The unique ID of the site that this reading belongs to
        Gateway_id int64 `json:"gateway_id,omitempty"`// The unique ID of the gateway that this reading belongs to
        Hub_id string `json:"hub_id,omitempty"` //  The unique ID of the hub that this reading belongs to
        Account_id int64 `json:"account_id,omitempty"`// The unique ID of the account that this reading belongs to
        Account_reference string `json:"account_reference,omitempty"`// The human-readable account number for the account this reading belongs to
        Credit_statuses []string `json:"credit_statuses,omitempty"`// An array of strings identifying what the credit status was for the account when this reading was generated
        Inst_demand float32 `json:"instantaneous_demand,omitempty"`// What the instantaneous demand was for the account at the time this reading was generated
        Inst_demand_formatted float32 `json:"instantaneous_demand_formatted,omitempty"`// The instantaneous demand formatted with the correct number of decimal places
        Inst_demand_units string `json:"instantaneous_demand_unit_of_measure,omitempty"`// The unit of measure for the instantaneous demand
        Reading_value float32 `json:"reading_value,omitempty"`// What the current summation was for the account at the time this reading was generated
        Reading_value_formatted float32 `json:reading_value_formatted,omitempty"`// The current summation formatted with the correct number of decimal places
        Reading_value_units string `json:"reading_value_unit_of_measure,omitempty"`// The unit of measure for the current summation
        Reading_date string `json:"reading_date,omitempty"`// The timestamp when this reading was generated, in ISO8601 format
        Balance float32 `json:"balance,omitempty"` // The credit balance when this reading was generated
        Currency string `json:"currency,omitempty"` // The currency for this account - typically GBP
        Received_at string `json:"received_at,omitempty"` //When this reading was received by the HES, in ISO8601 format
        Links map[string]map[string]string `json:"_links,omitempty"`
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

//HesSites implements sort.Interface
type HesSites []HesSite

func (ss HesSites) Len() int {
        return len(ss)
}

func (ss HesSites) Swap(i,j int) {
        ss[i], ss[j] = ss[j], ss[i]
}

func (ss HesSites) Less(i, j int) bool {
    return ss[i].Name < ss[j].Name
}

//HesSites implements sort.Interface
type HesGateways []HesGateway

func (gs HesGateways) Len() int {
        return len(gs)
}

func (gs HesGateways) Swap(i,j int) {
        gs[i], gs[j] = gs[j], gs[i]
}

func (gs HesGateways) Less(i, j int) bool {
    return gs[i].Name < gs[j].Name
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

