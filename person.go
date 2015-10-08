package basecamp

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email_address"`
	Admin     bool   `json:"admin"`
	Client    *Client
	AccountId int
}

func (p *Person) Events() ([]*Event, error) {
	url := fmt.Sprintf(baseURL, p.AccountId, fmt.Sprintf("people/%v/events.json?since=2012-03-24T11:00:00-06:00", p.Id))
	b, err := p.Client.get(url)
	if err != nil {
		return nil, err
	}
	var result []*Event
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	return result, nil
}
