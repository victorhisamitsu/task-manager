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
	r.Post("/{id}/add", handler.CreateNoteHandler)
	r.Put("/{noteID}/update", handler.ChangeNoteHandler)
	r.Delete("/{noteID}/delete", handler.DeleteNoteHandler)
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
	noteID, err := h.Service.CreateNote(ctx, id, bodyRequest.Content, bodyRequest.Order)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
		return
	}
	resposta["noteID"] = noteID
	httphandler.RespondSucess(resposta, w)
}

func (h Handler) ChangeNoteHandler(w http.ResponseWriter, r *http.Request) {

	bodyRequest := &models.NoteDto{}
	noteID := chi.URLParam(r, "noteID")
	resposta := map[string]any{"Sucess": true}
	ctx := context.Background()
	err := httphandler.ReadBody(r.Body, bodyRequest)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	resp, err := h.Service.ChangeNote(ctx, noteID, bodyRequest.Content, bodyRequest.Order)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
		return
	}
	//Responder
	resposta["Note"] = resp
	httphandler.RespondSucess(resposta, w)

}

func (h Handler) DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	//Ler Body
	noteID := chi.URLParam(r, "noteID")
	resposta := map[string]any{"Sucess": true}
	ctx := context.Background()

	//Executar minha service
	resp, err := h.Service.DeleteNote(ctx, noteID)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
		return
	}
	//Responder
	resposta["Delete"] = resp
	httphandler.RespondSucess(resposta, w)
}
