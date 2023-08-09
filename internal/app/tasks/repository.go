package tasks

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Hitsa/task-manager/internal/models"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type TasksRepository struct {
	DB *bun.DB
}

func NewRepositoryTasks(d *bun.DB) *TasksRepository {
	return &TasksRepository{
		DB: d,
	}
}

func (r *TasksRepository) NewTask(ctx context.Context, title string, description string, status string, dueData time.Time, important bool) (string, error) {

	// Criar task
	dateNow := time.Now()

	id := uuid.NewString()
	newTask := models.Task{
		ID:         id,
		Title:      title,
		Status:     status,
		DueDate:    dueData,
		Importante: important,
		CreatedAt:  dateNow,
	}

	res, err := r.DB.NewInsert().Model(&newTask).Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if res == nil {
		return "", nil
	}

	// Return
	fmt.Println(res)
	return id, nil

}

func (r *TasksRepository) GetAll(ctx context.Context) ([]models.Task, error) {
	listTasks := make([]models.Task, 0)
	_, err := r.DB.NewRaw("SELECT * FROM public.tasks").Exec(ctx, &listTasks)
	if err != nil {
		return nil, errors.New("nenhuma task cadastrada")
	}

	return listTasks, nil
}

func (r *TasksRepository) ChangeTask(ctx context.Context, id string, title string, description string, dueData time.Time, important bool) (*models.Task, error) {
	task := []models.Task{}

	// Consulta Db para ver se já existe task com mesmio titulo
	count, err := r.DB.NewSelect().Model(&task).Where("id=?", id).ScanAndCount(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, fmt.Errorf("task não encontrada")
	}

	//Alterar dados do DB
	dateNow := time.Now()

	tasks := models.Task{
		Title:      title,
		DueDate:    dueData,
		Importante: important,
		CreatedAt:  dateNow,
	}

	res, err := r.DB.NewUpdate().Model(&tasks).OmitZero().Where("id=?", id).Returning("*").Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if res == nil {
		return nil, errors.New("erro na busca de task")
	}

	// Return
	fmt.Println(res)
	return &tasks, nil

}

func (r *TasksRepository) GetTask(ctx context.Context, id string) (*models.Task, error) {
	task := models.Task{}
	_, err := r.DB.NewSelect().Model(&task).Where("id=?", id).Exec(ctx, &task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TasksRepository) DeleteTask(ctx context.Context, id string) (bool, error) {

	task := models.Task{}
	res, err := r.DB.NewDelete().Model(&task).Where("id=?", id).Exec(ctx)
	if err != nil {
		return false, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	//Verificar se houve modificações
	if count == 0 {
		return false, errors.New("nenhuma task encontrada")
	}
	return true, nil
}

func (r *TasksRepository) GetTasksFilter(ctx context.Context, query string) ([]models.Task, error) {
	listTasks := make([]models.Task, 0)
	err := r.DB.NewRaw("SELECT * FROM public.tasks WHERE ?", query).Scan(ctx, &listTasks)
	if err != nil {
		return nil, errors.New("nenhuma task cadastrada")
	}

	return listTasks, nil
}
