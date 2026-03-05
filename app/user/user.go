package user

type User struct {
	ID          uint64 `json:"id"`
	FullName    string `json:"fullname"`
	ChatID      int64  `json:"chat_id"`
	DateCreated int64  `json:"date_created"`
	DateUpdated int64  `json:"date_updated"`
}

func IsNil(user User) bool {
	return user.FullName == "" && user.ChatID == 0
}
