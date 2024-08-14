package usecase


import  (
	"github.com/mentesnot-2/refactoring_with_clean_architecture/Repository"	
	"github.com/mentesnot-2/refactoring_with_clean_architecture/Domain"

)



type TaskUsecase interface {
	CreateTask(task *domain.Task) error
	FindTaskById(id string) (*domain.Task, error)
	GetAllTask() ([]*domain.Task, error)
	UpdateTask(task *domain.Task) error
	DeleteTask(id string) error
}


type taskUsecaseImpl struct {
	taskRepo repository.TaskRepository
}

func NewTaskUsecase(taskRepo domain.TaskRepository) TaskUsecase {
	return &taskUsecaseImpl{taskRepo}
}

func (u *taskUsecaseImpl) CreateTask(task *domain.Task) error {
	return u.taskRepo.CreateTask(task)
}
func (u *taskUsecaseImpl) FindTaskById(id string) (*domain.Task, error) {
	return u.taskRepo.FindTaskById(id)
}

func (u *taskUsecaseImpl) GetAllTask() ([]*domain.Task, error) {
	return u.taskRepo.GetAllTask()
}

func (u *taskUsecaseImpl) UpdateTask(task *domain.Task) error {
	return u.taskRepo.UpdateTask(task)
}


func (u *taskUsecaseImpl) DeleteTask(id string) error {
	return u.taskRepo.DeleteTask(id)
}