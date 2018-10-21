package task

const ( //todo enums with relation type and profession type
	OWNER  = "Owner"
	LAWYER = "Lawyer"
)

type UserTaskRelation struct {
	Id         int
	UserId     int
	TaskId     int
	Relation   string
	Profession string
}

type UserTaskRelationRepository interface {
	Store(relation *UserTaskRelation) error
	FindById(id int) *UserTaskRelation
	FindOwnershipByTask(task *Task) *UserTaskRelation
}
