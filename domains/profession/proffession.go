package professionDomain

type ProfessionRepository interface {
	Store(user *Profession) error
	FindByUser(userId int) (*Profession, error)
}

type Profession struct {
	Id int

	UserId int

	/** name of profession ('CLIENT', 'LAWYER'...) */
	ProfessionType string

	/** id of specific profession (client, lawyer...) */
	ProfessionId int

	/** data of specific profession */
	data map[string]string
}
