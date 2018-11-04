package professionDomain

type LawyerRepository interface {
	Store(user *Lawyer) error
	FindById(id int) (*Lawyer, error)
}

type Lawyer struct {
	Id int
}
