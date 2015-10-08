package basecamp

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Href    string `json:"href"`
	Product string `json:"product"`
	Client  *Client
}

func (a *Account) Events() (events []Event, err error) {
  page := 1

  resultPage := a.getEvents(page)

  for i := 0; i < len(resultPage); i++ {
    resultPage[i].Client = a.Client
    events = append(events, resultPage[i])
  }

  for len(resultPage) == 50 {
    page += 1
    resultPage = a.getEvents(page)
    for i := 0; i < len(resultPage); i++ {
      resultPage[i].Client = a.Client
      events = append(events, resultPage[i])
    }
  }

  return events, nil
}

func (a *Account) getEvents(page int) (result []Event) {
  url := fmt.Sprintf(baseURL, a.Id, fmt.Sprintf("events.json?page=%v", page))

  b, err := a.Client.get(url)
  if err != nil {
    panic(err)
  }

  if err := json.Unmarshal(b, &result); err != nil {
    panic(err)
  }

  return result
}
