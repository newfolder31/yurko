package task

import (
	"yurko/domains/task"
	"yurko/usecases/communication"
)

type TaskInteractor struct {
	TaskRepository          task.TaskRepository
	RelationRepository      task.UserTaskRelationRepository
	CommunicationRepository communication.CommunicationRepository
}

func (interactor *TaskInteractor) CreateAnnounce(userId int, description string) (*task.Task, error) {
	//todo: validation // or not here?
	//todo: is userId is client? (wait for client/lawyer impl)
	//todo: error handling
	announce := task.Task{Description: description}
	err := interactor.TaskRepository.Store(&announce)

	relation := task.UserTaskRelation{UserId: userId, TaskId: announce.Id, Relation: task.OWNER, Profession: task.LAWYER}
	err := interactor.RelationRepository.Store(&relation)

	return &announce, nil
}

func (interactor *TaskInteractor) CreateRequest(userId int, description string, lawyerUserId int) (*task.Task, error) {
	request, err := interactor.CreateAnnounce(userId, description)
	err := interactor.AssignTask(request, lawyerUserId)
	return request, nil
}

func (interactor *TaskInteractor) AssignTask(task *task.Task, lawyerUserId int) error {
	//todo: is lawyerUserId a lawyer?
	ownershipRelation := interactor.RelationRepository.FindOwnershipByTask(task)
	_, err := interactor.CommunicationRepository.Store(ownershipRelation.Id, lawyerUserId)
	return err
}

func (interactor *TaskInteractor) TaskList(userId int) []*task.Task {
	return interactor.TaskRepository.FindByOwner(userId)
}

func (interactor *TaskInteractor) Task(taskId int) *task.Task {
	return interactor.TaskRepository.FindById(taskId)
}
