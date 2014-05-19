package twitter

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (c *Client) HomeTimeline() ([]Tweet, error) {
	response, err := c.get(
		c.apiUrl("/1.1/statuses/home_timeline.json"),
		map[string]string{},
	)
	if err != nil {
		return nil, err
	}

	return c.tweetsByResponse(response)
}

func (c *Client) MentionsTimeline() ([]Tweet, error) {
	response, err := c.get(
		c.apiUrl("/1.1/statuses/mentions_timeline.json"),
		map[string]string{},
	)
	if err != nil {
		return nil, err
	}

	return c.tweetsByResponse(response)
}

func (c *Client) UserTimeline(screenName string) ([]Tweet, error) {
	response, err := c.get(
		c.apiUrl("/1.1/statuses/user_timeline.json"),
		map[string]string{
			"screen_name": screenName,
		},
	)
	if err != nil {
		return nil, err
	}

	return c.tweetsByResponse(response)
}

func (c *Client) UpdateStatus(text string) error {
	_, err := c.post(
		c.apiUrl("/1.1/statuses/update.json"),
		map[string]string{
			"status": text,
		},
	)
	return err
}

func (c *Client) tweetsByResponse(response *http.Response) ([]Tweet, error) {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(bytes.NewReader(data))
	tweets := []Tweet{}
	decoder.Decode(&tweets)
	return tweets, nil
}
