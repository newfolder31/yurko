package task

type TaskRepository interface {
	Store(task *Task) error
	FindById(id int) *Task
	FindByOwner(ownerId int) []*Task
}

type Task struct {
	Id          int
	Description string
}
