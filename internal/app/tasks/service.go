package tasks

import (
	"context"
	"errors"
	"time"

	"github.com/Hitsa/task-manager/internal/models"
)

type TasksService struct {
	repository *TasksRepository
}

func NewTasksService(r *TasksRepository) *TasksService {
	tasks := TasksService{r}
	return &tasks
}

func (s TasksService) CreateTask(ctx context.Context, title string, description string, status string, dueData time.Time, important bool) (string, error) {
	if title == "" {
		return "nil", errors.New("titulo não pode ser")
	}
	id, err := s.repository.NewTask(ctx, title, description, status, dueData, important)
	if err != nil {
		return "nil", err
	}
	return id, nil
}

func (s TasksService) GetAll(ctx context.Context, filter string) ([]models.Task, error) {

	switch filter {
	case "all":
		dateNow := time.Now().Format("02/01/2006")
		query := "DueDate = " + dateNow
		listTask, err := s.repository.GetTasksFilter(ctx, query)
		if err != nil {
			return nil, err
		}
		return listTask, nil
	case "today":
		dateNow := time.Now().Format("02/01/2006")
		listTask, err := s.repository.GetTasksFilter(ctx, dateNow)
		if err != nil {
			return nil, err
		}
		return listTask, nil
	case "month":
		dateNow := time.Now().Format("02/01/2006")
		afterMonth := time.Now().Add(30 * 24 * time.Hour).Format("02/01/2006")
		query := "DueDate >= " + dateNow + " AND DueDate <= " + afterMonth
		listTask, err := s.repository.GetTasksFilter(ctx, query)
		if err != nil {
			return nil, err
		}
		return listTask, nil
	case "week":
		dateNow := time.Now().Format("02/01/2006")
		afterMonth := time.Now().Add(7 * 24 * time.Hour).Format("02/01/2006")
		query := "DueDate >= " + dateNow + " AND DueDate <= " + afterMonth
		listTask, err := s.repository.GetTasksFilter(ctx, query)
		if err != nil {
			return nil, err
		}
		return listTask, nil
	case "done":
		dateNow := time.Now().Format("02/01/2006")
		query := "DueDate <= " + dateNow
		listTask, err := s.repository.GetTasksFilter(ctx, query)
		if err != nil {
			return nil, err
		}
		return listTask, nil
	default:
		listTask, err := s.repository.GetAll(ctx)
		if err != nil {
			return nil, err
		}
		return listTask, nil
	}
}

func (s TasksService) ChangeTask(ctx context.Context, id string, title string, description string, dueData time.Time, important bool) (*models.Task, error) {
	if id == "" {
		return nil, errors.New("id nulo")
	}
	resp, err := s.repository.ChangeTask(ctx, id, title, description, dueData, important)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s TasksService) GetTask(ctx context.Context, id string) (*models.Task, error) {
	task, err := s.repository.GetTask(ctx, id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s TasksService) DeleteTask(ctx context.Context, id string) (bool, error) {
	valid, err := s.repository.DeleteTask(ctx, id)
	if err != nil {
		return false, err
	}
	return valid, nil
}
