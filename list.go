package twitter

type List struct {
	FullName    string `json:"full_name"`
	Name        string
	MemberCount int `json:"member_count"`
	Description string
	User        *User
}
