package domain

type UserRepository interface {
	CreateUser(name string) (User, error)
	GetUserByID(id int) (User, error)
}
