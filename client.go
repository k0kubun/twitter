package twitter

import (
	"bytes"
	"encoding/json"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	twitterApiUrl = "https://api.twitter.com"

	requestTokenPath   = "/oauth/request_token"
	authorizeTokenPath = "/oauth/authorize"
	accessTokenPath    = "/oauth/access_token"

	homeTimelinePath = "/1.1/statuses/home_timeline.json"
)

type Client struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string

	cons  *oauth.Consumer
	token *oauth.AccessToken
}

func (c *Client) HomeTimeline() []Tweet {
	response, err := c.get(
		c.apiUrl(homeTimelinePath),
		map[string]string{},
	)
	if err != nil {
		log.Println(err)
		return nil
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(bytes.NewReader(data))
	tweets := []Tweet{}
	decoder.Decode(&tweets)
	return tweets
}

func (c *Client) get(requestUrl string, params map[string]string) (*http.Response, error) {
	return c.consumer().Get(requestUrl, params, c.accessToken())
}

func (c *Client) post(requestUrl string, params map[string]string) (*http.Response, error) {
	return c.consumer().Post(requestUrl, params, c.accessToken())
}

func (c *Client) consumer() *oauth.Consumer {
	if c.cons != nil {
		return c.cons
	}

	c.cons = oauth.NewConsumer(
		c.ConsumerKey,
		c.ConsumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   c.apiUrl(requestTokenPath),
			AuthorizeTokenUrl: c.apiUrl(authorizeTokenPath),
			AccessTokenUrl:    c.apiUrl(accessTokenPath),
		},
	)
	return c.cons
}

func (c *Client) accessToken() *oauth.AccessToken {
	if c.token != nil {
		return c.token
	}

	c.token = &oauth.AccessToken{
		Token:  c.AccessToken,
		Secret: c.AccessTokenSecret,
	}
	return c.token
}

func (c *Client) apiUrl(apiPath string) string {
	return twitterApiUrl + apiPath
}
