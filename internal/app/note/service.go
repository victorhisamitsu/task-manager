package note

import (
	"context"

	"github.com/Hitsa/task-manager/internal/models"
)

type NoteService struct {
	repository *NoteRepository
}

func NewNoteService(r *NoteRepository) *NoteService {
	notes := NoteService{r}
	return &notes
}

// Criar uma nova nota
func (s NoteService) CreateNote(ctx context.Context, id string, content string, order string) (string, error) {
	noteID, err := s.repository.NewNote(ctx, id, content, order)
	if err != nil {
		return "", err
	}
	return noteID, nil
}

// Alterar uma nota
func (s NoteService) ChangeNote(ctx context.Context, id string, content string, order string) (*models.Note, error) {
	resp, err := s.repository.ChangeNote(ctx, id, content, order)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Deletar uma nota
func (s NoteService) DeleteNote(ctx context.Context, id string) (bool, error) {
	valid, err := s.repository.DeleteNote(ctx, id)
	if err != nil {
		return false, err
	}
	return valid, nil
}

// Buscar notas que pertencem a uma task
func (s NoteService) GetNoteByTaskID(ctx context.Context, taskID string) ([]models.Note, error) {
	resp, err := s.repository.GetNoteByTaskID(ctx, taskID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
