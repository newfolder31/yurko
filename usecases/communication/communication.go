package communication

type CommunicationRepository interface {
	Store(relationId int, lawyerUserId int) (int, error)
	FindById(id int) (int, int)
}
