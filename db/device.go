package db

type Device struct {
	Model
	UserID string `json:"user_id"`
}
