package twitter

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (c *Client) ReplyStatus(text string, tweetId int64) error {
	_, err := c.post(
		c.apiUrl("/1.1/statuses/update.json"),
		map[string]string{
			"status":                text,
			"in_reply_to_status_id": fmt.Sprintf("%d", tweetId),
		},
	)
	return err
}

func (c *Client) Favorite(tweetId int64) error {
	_, err := c.post(
		c.apiUrl("/1.1/favorites/create.json"),
		map[string]string{
			"id": fmt.Sprintf("%d", tweetId),
		},
	)
	return err
}

func (c *Client) Retweet(tweetId int64) error {
	_, err := c.post(
		c.apiUrl("/1.1/statuses/retweet/%d.json", tweetId),
		map[string]string{},
	)
	return err
}

func (c *Client) Destroy(tweetId int64) error {
	_, err := c.post(
		c.apiUrl("/1.1/statuses/destroy/%d.json", tweetId),
		map[string]string{},
	)
	return err
}

func (c *Client) Lists() ([]List, error) {
	response, err := c.get(
		c.apiUrl("/1.1/lists/list.json"),
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
	lists := []List{}
	decoder.Decode(&lists)
	return lists, nil
}

func (c *Client) ListTimeline(ownerScreenName string, slug string) ([]Tweet, error) {
	response, err := c.get(
		c.apiUrl("/1.1/lists/statuses.json"),
		map[string]string{
			"owner_screen_name": ownerScreenName,
			"slug":              slug,
		},
	)
	if err != nil {
		return nil, err
	}

	return c.tweetsByResponse(response)
}

func (c *Client) AddToList(ownerScreenName string, screenName string, slug string) error {
	_, err := c.post(
		c.apiUrl("/1.1/lists/members/create.json"),
		map[string]string{
			"owner_screen_name": ownerScreenName,
			"screen_name":       screenName,
			"slug":              slug,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Search(query string) ([]Tweet, error) {
	response, err := c.get(
		c.apiUrl("/1.1/search/tweets.json"),
		map[string]string{
			"q": query,
		},
	)
	if err != nil {
		return nil, err
	}

	decoder := c.jsonDecoder(response)
	result := SearchResult{}
	decoder.Decode(&result)
	return result.Statuses, nil
}

func (c *Client) tweetsByResponse(response *http.Response) ([]Tweet, error) {
	decoder := c.jsonDecoder(response)
	tweets := []Tweet{}
	decoder.Decode(&tweets)
	return tweets, nil
}

func (c *Client) jsonDecoder(response *http.Response) *json.Decoder {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return json.NewDecoder(bytes.NewReader(data))
}
