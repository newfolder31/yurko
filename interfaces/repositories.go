package interfaces

import "yurko/domains/task"

var index = 0

type MemoStorage struct {
	Tasks          map[int]*task.Task
	Relations      map[int]*task.UserTaskRelation
	Communications map[int]int // relation, lawyer
}

func NewMemoStorage() *MemoStorage {
	storage := MemoStorage{}
	storage.Tasks = make(map[int]*task.Task)
	storage.Relations = make(map[int]*task.UserTaskRelation)
	storage.Communications = make(map[int]int)

	return &storage
}

type TaskInMemoRepo struct {
	Storage *MemoStorage
}

func (repo *TaskInMemoRepo) Store(task *task.Task) error {
	defer func() { index++ }()

	task.Id = index
	repo.Storage.Tasks[index] = task
	return nil
}

func (repo *TaskInMemoRepo) FindById(id int) *task.Task {
	task, present := repo.Storage.Tasks[id]
	if present {
		return task
	} else {
		return nil
	}
}

func (repo *TaskInMemoRepo) FindByOwner(id int) []*task.Task {
	tasks := make([]*task.Task, 0)
	ids := make([]int, 0)
	for k, v := range repo.Storage.Relations {
		if v.UserId == id {
			ids = append(ids, k)
		}
	}

	for _, v := range repo.Storage.Tasks {
		for _, id := range ids {
			if v.Id == id {
				tasks = append(tasks, v)
			}
		}
	}
	return tasks
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type RelationInMemoRepo struct {
	Storage *MemoStorage
}

func (repo *RelationInMemoRepo) Store(relation *task.UserTaskRelation) error {
	defer func() { index++ }()

	relation.Id = index
	repo.Storage.Relations[index] = relation
	return nil
}

func (repo *RelationInMemoRepo) FindById(id int) *task.UserTaskRelation {
	relation, present := repo.Storage.Relations[id]
	if present {
		return relation
	} else {
		return nil
	}
}

func (repo *RelationInMemoRepo) FindOwnershipByTask(Task *task.Task) *task.UserTaskRelation {
	for _, v := range repo.Storage.Relations {
		if v.TaskId == Task.Id && v.Relation == task.OWNER {
			return v
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type CommunicationInMemoRepo struct {
	Storage *MemoStorage
}

func (repo *CommunicationInMemoRepo) Store(relationId int, lawyerUserId int) (int, error) {
	repo.Storage.Communications[relationId] = lawyerUserId
	return 0, nil //todo
}

func (repo *CommunicationInMemoRepo) FindById(id int) (int, int) {
	return 0, 0
	//todo
}
