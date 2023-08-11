package note

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Hitsa/task-manager/internal/models"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type NoteRepository struct {
	DB *bun.DB
}

func NewRepositoryNote(d *bun.DB) *NoteRepository {
	return &NoteRepository{
		DB: d,
	}
}

// Nova nota
func (r *NoteRepository) NewNote(ctx context.Context, id string, content string, order string) (string, error) {

	//Verificar se existe task com esse id
	task := models.Task{}
	_, err := r.DB.NewSelect().Model(&task).Where("id=?", id).Exec(context.Background(), &task)
	if err != nil {
		return "", errors.New("task não encontrada")
	}
	// Adicionar Note
	dateNow := time.Now()
	noteID := uuid.NewString()

	newNote := models.Note{
		ID:        noteID,
		Content:   content,
		CreatedAt: dateNow,
		Order:     order,
		TaskId:    id,
	}
	res, err := r.DB.NewInsert().Model(&newNote).Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if res == nil {
		return "", nil
	}
	return noteID, nil
}

// Alterar nota
func (r *NoteRepository) ChangeNote(ctx context.Context, id string, content string, order string) (*models.Note, error) {
	notes := []models.Note{}

	// Consulta Db para ver se já existe task com mesmio titulo
	count, err := r.DB.NewSelect().Model(&notes).Where("id=?", id).ScanAndCount(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, fmt.Errorf("nota não encontrada")
	}

	//Alterar dados do DB
	note := models.Note{
		Content: content,
		Order:   order,
	}

	res, err := r.DB.NewUpdate().Model(&note).OmitZero().Where("id=?", id).Returning("*").Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if res == nil {
		return nil, errors.New("erro na busca da nota")
	}

	// Return
	fmt.Println(res)
	return &note, nil
}

// Deletar nota
func (r *NoteRepository) DeleteNote(ctx context.Context, id string) (bool, error) {

	note := models.Note{}
	res, err := r.DB.NewDelete().Model(&note).Where("id=?", id).Exec(ctx)
	if err != nil {
		return false, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	//Verificar se houve modificações
	if count == 0 {
		return false, errors.New("nenhuma nota encontrada")
	}
	return true, nil
}

// Buscar todas as notas que pertecem a uma task
func (r *NoteRepository) GetNoteByTaskID(ctx context.Context, taskID string) ([]models.Note, error) {
	notes := []models.Note{}

	// Consulta Db para ver se já existe task com mesmio titulo
	_, err := r.DB.NewSelect().Model(&notes).Where("TaskId=?", taskID).Exec(ctx, &notes)
	if err != nil {
		return nil, err
	}
	// Return
	return notes, nil
}
