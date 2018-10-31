package userUsecases

type UserRepository interface {
	Store(user *User) error
	FindById(id int) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByEmailAndPassword(email, password string) (*User, error)
}

type User struct {
	Id       int
	Email    string
	Password string

	IsAdmin bool

	IsActive bool

	FirstName   string
	LastName    string
	FathersName string
}
