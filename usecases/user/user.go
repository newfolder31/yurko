package user

type UserRepository interface {
	Store(user *User) error
	FindById(id int) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByEmailAndPassword(email, password string) (*User, error)
}

type AddressRepository interface {
	Store(address *Address) error
	FindById(id int) (*Address, error)
}

//table_name: usr
type User struct {
	Id       int
	Email    string
	Password string

	IsAdmin  bool `db:"is_admin"`
	IsActive bool `db:"is_active"`

	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	FathersName string `db:"fathers_name"`

	Phone string `json:"phone"`

	AddressId int `db:"address_id"`
}

type Address struct {
	Id int

	Building string `db:"building"`
	Street   string `db:"street"`
	City     string `db:"city"`
}
