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

func (c *Client) HomeTimeline() ([]Tweet, error) {
	response, err := c.get(
		c.apiUrl(homeTimelinePath),
		map[string]string{},
	)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(bytes.NewReader(data))
	tweets := []Tweet{}
	decoder.Decode(&tweets)
	return tweets, nil
}
