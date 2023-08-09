package note

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Hitsa/task-manager/internal/httphandler"
	"github.com/Hitsa/task-manager/internal/models"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Service *NoteService
}

func NewNotesHandler(service *NoteService) *chi.Mux {
	handler := Handler{
		Service: service,
	}
	r := chi.NewRouter()
	r.Post("/", handler.BodyApiHandlerNote)
	r.Post("/{id}/notes/add", handler.CreateNoteHandler)
	r.Put("/{id}/notes/update", handler.ChangeNoteHandler)
	r.Delete("/{id}/notes/delete", handler.DeleteNoteHandler)
	return r
}

func (h Handler) BodyApiHandlerNote(w http.ResponseWriter, r *http.Request) {
	resposta := map[string]any{}
	body, _ := io.ReadAll(r.Body)
	var minhaVariavel map[string]string
	json.Unmarshal(body, &minhaVariavel)
	resposta["resposta"] = minhaVariavel
	variavelJson, _ := json.Marshal(minhaVariavel)

	w.Write(variavelJson)
}

func (h Handler) CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	bodyRequest := &models.NoteDto{}
	id := chi.URLParam(r, "id")
	resposta := map[string]any{"Sucess": true}
	ctx := context.Background()
	err := httphandler.ReadBody(r.Body, bodyRequest)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	err = h.Service.CreateNote(ctx, id, bodyRequest.Content, bodyRequest.Order)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
	}
}

func (h Handler) ChangeNoteHandler(w http.ResponseWriter, r *http.Request) {

	bodyRequest := &models.NoteDto{}
	id := chi.URLParam(r, "id")
	resposta := map[string]any{"Sucess": true}
	ctx := context.Background()
	err := httphandler.ReadBody(r.Body, bodyRequest)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	resp, err := h.Service.ChangeNote(ctx, id, bodyRequest.Content, bodyRequest.Order)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
	}
	//Responder
	resposta["Note"] = resp
	httphandler.RespondSucess(resposta, w)

}

func (h Handler) DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	//Ler Body
	bodyRequest := &models.TaskDto{}
	id := chi.URLParam(r, "id")
	resposta := map[string]any{"Sucess": true}
	ctx := context.Background()
	err := httphandler.ReadBody(r.Body, bodyRequest)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	resp, err := h.Service.DeleteNote(ctx, id)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
	}
	//Responder
	resposta["Note"] = resp
	httphandler.RespondSucess(resposta, w)
}
