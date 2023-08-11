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

// Criar uma nova task
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

// Buscar tasks com filtro de dias ou ALL
func (s TasksService) GetAll(ctx context.Context, filter string) ([]models.Task, error) {

	switch filter {

	//Todas as tasks em um intervalo até o fim do dia
	case "today":
		listTask, err := s.getFutureDaysAndGetTasks(ctx, 1)
		if err != nil {
			return nil, err
		}
		return listTask, nil

	//Todas as tasks em um intervalo de 30 dias
	case "month":
		listTask, err := s.getFutureDaysAndGetTasks(ctx, 30)
		if err != nil {
			return nil, err
		}
		return listTask, nil

	//Todas as tasks em um intervalo de 7 dias
	case "week":
		listTask, err := s.getFutureDaysAndGetTasks(ctx, 7)
		if err != nil {
			return nil, err
		}
		return listTask, nil

	//Todas as tasks concluidas
	case "done":
		dateNow := getDateToday()
		listTask, err := s.repository.GetTasksFilter(ctx, dateNow)
		if err != nil {
			return nil, err
		}
		return listTask, nil

	//ALL buscar todas as datas quando default, funciona para função ALL
	default:
		listTask, err := s.repository.GetAll(ctx)
		if err != nil {
			return nil, err
		}
		return listTask, nil
	}
}

// Muda qualquer dado da task, mas não pode mudar status
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

// Buscar task pelo ID
func (s TasksService) GetTask(ctx context.Context, id string) (*models.Task, error) {
	task, err := s.repository.GetTask(ctx, id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// Apagar uma task
func (s TasksService) DeleteTask(ctx context.Context, id string) (bool, error) {
	valid, err := s.repository.DeleteTask(ctx, id)
	if err != nil {
		return false, err
	}
	return valid, nil
}

// Muda o status da task
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

// Busca tasks com notas, caso não tenha notas retorna lista vazia
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

// Funcão que busca tasks em um determinado intervalo de tempo e retorna a lista de tasks
func (s TasksService) getFutureDaysAndGetTasks(ctx context.Context, days int) ([]models.Task, error) {
	currentDay := getDateToday()
	afterDays := currentDay.Add(time.Duration(days) * 24 * time.Hour)
	listTask, err := s.repository.GetTasksFilterBetween(ctx, currentDay, afterDays)
	if err != nil {
		return nil, err
	}
	return *listTask, nil
}

// Busca data de hoje desde o começo do dia
func getDateToday() time.Time {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return startOfDay
}
