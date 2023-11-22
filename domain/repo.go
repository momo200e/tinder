package domain

type Repository interface {
	AddUser(user User)
	DeleteUserByName(name string)
	GetUserByName(name string) *User
	UpdateUserRemainDatesDecrByName(name string) (uint8, error)
	FindUsersByGenderAndHeight(gender Gender, height uint8, isFindGreater bool, count int) []*User
}
