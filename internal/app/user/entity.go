package user

type User struct {
	UserId   int64  `json:"user_id"`
	UserName string `db:"user_name"`
}

func NewUser(userId int64, userName string) User {
	var user User
	user.UserId = userId
	user.UserName = userName
	return user
}
