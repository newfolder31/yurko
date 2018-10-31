package profession


type ProfessionRepository interface {
	Store(user *Profession) error
	FindByUser(id int) (*Profession, error)
}

type Profession struct {

	Id int

	UserId int
	ProfessionType string
	ProfessionId int

	data map[string]string
}