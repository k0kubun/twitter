package twitter

type Tweet struct {
	Id              int64
	User            *User
	Source          string
	Text            string
	CreatedAt       string `json:"created_at"`
	Retweeted       bool
	RetweetedStatus *Tweet `json:"retweeted_status"`
}
