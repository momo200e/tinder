package domain

type Service interface {
	AddSinglePersonAndMatch(newUser *User) ([]*User, *ErrorFormat)
	QuerySinglePerson(userName string, number int) ([]*User, *ErrorFormat)
	RemoveSinglePerson(userName string) *ErrorFormat
}
