# Twitter

Tiny twitter client library for golang

## Usage

```go
import "github.com/k0kubun/twitter"

func main() {
	client := &twitter.Client{
		ConsumerKey:       "CONSUMER_KEY",
		ConsumerSecret:    "CONSUMER_SECRET",
		AccessToken:       "ACCESS_TOKEN",
		AccessTokenSecret: "ACCESS_TOKEN_SECRET",
	}

	tweets, _ := client.HomeTimeline()
	for _, tweet := range tweets {
		fmt.Printf("%s: %s\n", tweet.User.ScreenName, tweet.Text)
	}

	client.UpdateStatus("Tweet test")
}
```

## License

MIT License
