package db

type User struct {
	Model
	Nickname string   `json:"nickname"`
	Avatar   string   `json:"avatar"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Salt     string   `json:"-"`
	OpenId   string   `json:"openid"`
	UnionID  string   `json:"unionid"`
	Devices  []Device `json:"devices"`
	Tokens   []Token  `json:"tokens"`
	Scripts  []Script `json:"scripts"`
}
