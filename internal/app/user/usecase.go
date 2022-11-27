package user

func NewUser(userId int64, userName string) User {
	var user User
	user.UserId = userId
	user.UserName = userName
	return user
}
