package communication

type Communication struct {
	Id           int
	RelationId   int
	LawyerUserId int
}

type CommunicationRepository interface {
	Store(*Communication) error
	FindById(id int) *Communication
}
