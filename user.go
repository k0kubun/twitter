package twitter

type User struct {
	Id         int64
	ScreenName string `json:"screen_name"`
	Protected  bool
}
