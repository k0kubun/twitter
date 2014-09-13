package twitter

type List struct {
	Slug        string `json:"slug"`
	FullName    string `json:"full_name"`
	Name        string `json:"name"`
	IdStr       string `json:"id_str"`
	MemberCount int    `json:"member_count"`
	Description string `json:"description"`
	User        *User
}
