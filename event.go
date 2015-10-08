package basecamp

import(
  "time"
)

type Event struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Private   bool      `json:"private"`
	Action    string    `json:"action"`
	Target    string    `json:"target"`
	Eventable Eventable `json:"eventable"`
	Creator   Creator   `json:"creator"`
  Bucket    Bucket    `json:"bucket"`
	Summary   string    `json:"summary"`
	Url       string    `json:"url"`
	HtmlUrl   string    `json:"html_url"`
	Client    *Client
}

type Eventable struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	Url    string `json:"url"`
	AppUrl string `json:"app_url"`
}

type Creator struct {
  Attachments       string `json:"attachments"`
  Excerpt           string `json:"excerpt"`
  RawExcerpt        string `json:"raw_excerpt"`
  Id                int    `json:"id"`
  Name              string `json:"name"`
  AvatarUrl         string `json:"avatar_url"`
  FullsizeAvatarUrl string `json:"fullsize_avatar_url"`
}

type Bucket struct {
	Type   string `json:"type"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Color  string    `json:"color"`
	Url    string `json:"url"`
	AppUrl string `json:"app_url"`
}
