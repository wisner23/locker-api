package models

type Users struct {
	UserId int `json:"user_id" orm:"auto"`
	Email string `json:"email" orm:"size(128)"`
	Username string `json:"	" orm:"size(32)"`
	Phone string `json:"phone" orm:"size(15)"`
}