package db

type Script struct {
	Model
	UserID string `json:"user_id"`
	//js或者lua
	Type     string `json:"type"`
	Filename string `json:"filename"`
	Content  string `json:"content"`
}
