package db

type Token struct {
	Model
	UserID string `json:"user_id"`
}
