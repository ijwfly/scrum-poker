package user

type Repository interface {
	GetUser(userId int64) (User, error)
}
