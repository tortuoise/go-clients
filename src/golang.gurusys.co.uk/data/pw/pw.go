package pw

import (
        "fmt"
        "net/http"
)

type ApiKey struct {
	Api_key          string `json:"api_key,omitempty"`
}

type Account struct {
        Id int64 `json:"id,omitempty"`
        Name string `json:"name,omitempty"`
}

type User struct {
        Id int64 `json:"id,omitempty"`
        Name string `json:"name,omitempty"`
        Email string `json:"email,omitempty"`
}

type Address struct {
        Street string `json:"street,omitempty"`
        City string `json:"city,omitempty"`
        State string `json:"state,omitempty"`
        Postal_code string `json:"postal_code,omitempty"`
        Country string `json:"country,omitempty"`
}

func (a *Address) String() string {
	return fmt.Sprintf(" %10d %10s %10s %10s %10s \n", a.Street, a.City, a.State, a.Postal_code, a.Country)
}

type Email map[string]string
type Phone map[string]string

type Lead struct {
        Id int64 `json:"id,omitempty"`
        Name string `json:"name,omitempty"`
        LEmail Email `json:"email,omitempty"`
        LAddress Address `json:"address,omitempty"`
        Company string `json:"company_name,omitempty"`
        Details string `json:"details,omitempty"`
        Assignee_id int64 `json:"assignee_id,omitempty"`
        Customer_source_id int64 `json:"customer_source_id,omitempty"`
}

func (l *Lead) String() string {
        if _,ok := l.LEmail["email"]; !ok {
                return fmt.Sprintf(" %10d %10s %10s %10s %10s \n", l.Id, l.Name, l.LAddress.City, l.Company, l.Details/*, l.Assignee_id, l.Customer_source_id*/)
        } else {
	        return fmt.Sprintf(" %10d %10s %10s %10s %10s %10s \n", l.Id, l.Name, l.LEmail["email"], l.LAddress.City, l.Company, l.Details/*, l.Assignee_id, l.Customer_source_id*/)
        }
}

type Leads []Lead

func (l *Leads) String() string {
	return fmt.Sprintf(" %10s %10s %10s %10s %10s %10s %10s \n", "Id", "Name", "Email", "Address", "Company", "Details", "Assignee_id", "Customer_source_id")
}

type Person struct {
        Id int64 `json:"id,omitempty"`
        Name string `json:"name,omitempty"`
        Emails []Email `json:"emails,omitempty"`
        Phones []Phone `json:"phone_numbers,omitempty"`
        LAddress Address `json:"address,omitempty"`
        Company_id string `json:"company_id,omitempty"`
        Company_name string `json:"company_name,omitempty"`
        Details string `json:"details,omitempty"`
        Assignee_id int64 `json:"assignee_id,omitempty"`
        Contact_type_id int64 `json:"contact_type_id,omitempty"`
}
func (p *Person) String() string {
        return fmt.Sprintf(" %10d %10s %10s %10s %10s %10d \n", p.Id, p.Name, p.LAddress.City, p.Company_name, p.Details, p.Assignee_id, /*l.Customer_source_id*/)
}
type People []Person

func (p *People) String() string {
	return fmt.Sprintf(" %10s %10s %10s %10s %10s %10s \n", "Id", "Name", "Address", "Company", "Details", "Assignee_id")
}

type Company struct {
        Id int64 `json:"id,omitempty"`
        Name string `json:"name,omitempty"`
        Emails []Email `json:"emails,omitempty"`
        Phones []Phone `json:"phone_numbers,omitempty"`
        LAddress Address `json:"address,omitempty"`
        Email_domain string `json:"email_domain,omitempty"`
        Details string `json:"details,omitempty"`
        Assignee_id int64 `json:"assignee_id,omitempty"`
        Contact_type_id int64 `json:"contact_type_id,omitempty"`
}
func (c *Company) String() string {
        return fmt.Sprintf(" %10d %10s %10s %10s %10s %10d \n", c.Id, c.Name, c.LAddress.City, c.Email_domain, c.Details, c.Assignee_id, /*l.Customer_source_id*/)
}

type Companies []Company

func (c *Companies) String() string {
	return fmt.Sprintf(" %10s %10s %10s %10s %10s %10s \n", "Id", "Name", "Address", "Domain", "Details", "Assignee_id")
}

type Opportunity struct {
        Id int64 `json:"id,omitempty"`
        Name string `json:"name,omitempty"`
        Company_id string `json:"company_id,omitempty"`
        Company_name string `json:"company_name,omitempty"`
        Details string `json:"details,omitempty"`
        Assignee_id int64 `json:"assignee_id,omitempty"`
        Close_date string `json:"close_date,omitempty"`
        Monetary_value int64 `json:"monetary_value,omitempty"`
}
func (o *Opportunity) String() string {
        return fmt.Sprintf(" %10d %10s %10s %10s %10d %10s %10d \n", o.Id, o.Name, o.Company_name, o.Close_date, o.Monetary_value, o.Details, o.Assignee_id, /*l.Customer_source_id*/)
}

type Opportunities []Opportunity

func (o *Opportunities) String() string {
	return fmt.Sprintf(" %10s %10s %10s %10s %10s %10s %10s \n", "Id", "Name", "Company", "Close_date", "Monetary_value", "Details", "Assignee")
}

type Custom_field map[int64]interface{}

type Project struct {
        Id int64 `json:"id,omitempty"`
        Name string `json:"name,omitempty"`
        Assignee_id int64 `json:"assignee_id,omitempty"`
        Details string `json:"details,omitempty"`
        Status string `json:"status,omitempty"`
        Date_created int64 `json:"date_created,omitempty"`
        Date_modified int64 `json:"date_modified,omitempty"`
        Custom_fields Custom_field `json:"custom_fields,omitempty"`
}

type Projects []Project

func (p *Projects) String() string {
	return fmt.Sprintf(" %10s %10s %10s %10s %10s %10s %10s \n", "Id", "Name", "Date_created", "Date_modified", "Status", "Details", "Assignee")
}

type Task struct {
        Id int64 `json:"id,omitempty"`
        Name string `json:"name,omitempty"`
        Assignee_id int64 `json:"assignee_id,omitempty"`
        Details string `json:"details,omitempty"`
        Due_date int64 `json:"due_date,omitempty"`
        Reminder_date int64 `json:"reminder_date,omitempty"`
        Completed_date int64 `json:"completed_date,omitempty"`
        Priority string `json:"priority,omitempty"`
        Status string `json:"status,omitempty"`
        Date_created int64 `json:"date_created,omitempty"`
        Date_modified int64 `json:"date_modified,omitempty"`
        Custom_fields Custom_field `json:"custom_fields,omitempty"`
}
type Tasks []Task

func (t *Tasks) String() string {
	return fmt.Sprintf(" %10s %10s %10s %10s %10s %10s %10s \n", "Id", "Name", "Due_date", "Completed_date", "Status", "Details", "Assignee")
}

type Activity struct {

        Id int64 `json:"id,omitempty"`
        Type string `json:"type,omitempty"`
        Details string `json:"details,omitempty"`
        User_id int64 `json:"user_id,omitempty"`
        Activity_date int64 `json:"activity_date,omitempty"`
}
type Activities []Activity

func (a *Activities) String() string {
	return fmt.Sprintf(" %10s %10s %10s %10s %10s \n", "Id", "Type", "User_id", "Activity_date", "Details")
}

//Leads implements sort.Interface
func (ls Leads) Len() int {
        return len(ls)
}

func (ls Leads) Swap(i,j int) {
        ls[i], ls[j] = ls[j], ls[i]
}

func (ls Leads) Less(i, j int) bool {
    return ls[i].Name < ls[j].Name
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
	req.Header.Set("X-PW-AccessToken", accesstoken)
	req.Header.Set("X-PW-Application", "developer_api")
	req.Header.Set("X-PW-UserEmail", "prachee.peshkar@gurusystems.com")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
}
