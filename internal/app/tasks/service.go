package tasks

import (
	"context"
	"errors"
	"time"

	"github.com/Hitsa/task-manager/internal/app/note"
	"github.com/Hitsa/task-manager/internal/models"
)

type TasksService struct {
	repository *TasksRepository
	note       *note.NoteService
}

func NewTasksService(r *TasksRepository, n *note.NoteService) *TasksService {
	tasks := TasksService{r, n}
	return &tasks
}

func (s TasksService) CreateTask(ctx context.Context, title string, description string, status string, dueData *time.Time, important bool) (string, error) {
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

	case "today":
		now := time.Now()
		startOfDay := getDateToday()
		endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
		listTask, err := s.repository.GetTasksFilterBetween(ctx, startOfDay, endOfDay)
		if err != nil {
			return nil, err
		}
		if *listTask == nil {
			return nil, errors.New("nenhuma task encontrada")
		}
		return *listTask, nil

	case "month":
		startOfDay := getDateToday()
		afterMonth := time.Now().Add(30 * 24 * time.Hour)
		listTask, err := s.repository.GetTasksFilterBetween(ctx, startOfDay, afterMonth)
		if err != nil {
			return nil, err
		}
		if *listTask == nil {
			return nil, errors.New("nenhuma task encontrada")
		}
		return *listTask, nil

	case "week":
		startOfDay := getDateToday()
		afterWeek := time.Now().Add(7 * 24 * time.Hour)
		listTask, err := s.repository.GetTasksFilterBetween(ctx, startOfDay, afterWeek)
		if err != nil {
			return nil, err
		}
		if *listTask == nil {
			return nil, errors.New("nenhuma task encontrada")
		}

		return *listTask, nil

	case "done":
		dateNow := getDateToday()
		listTask, err := s.repository.GetTasksFilter(ctx, dateNow)
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

func (s TasksService) ChangeTask(ctx context.Context, id string, title string, description string, dueData *time.Time, important bool) (*models.Task, error) {
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

func (s TasksService) ChangeStatus(ctx context.Context, id string, status string) (*models.Task, error) {
	if id == "" {
		return nil, errors.New("id nulo")
	}
	resp, err := s.repository.ChangeStatus(ctx, id, status)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s TasksService) GetTaskWithNote(ctx context.Context, id string) (*models.Task, error) {
	task, err := s.repository.GetTasksWithNote(ctx, id)
	if err != nil {
		return nil, err
	}
	litasNote, err := s.note.GetNoteByTaskID(ctx, id)
	if err != nil {
		return nil, err
	}
	task.Notes = litasNote
	return task, nil
}

func getDateToday() time.Time {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return startOfDay
}
