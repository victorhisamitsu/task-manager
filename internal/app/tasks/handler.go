package tasks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Hitsa/task-manager/internal/httphandler"
	"github.com/Hitsa/task-manager/internal/models"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Service *TasksService
}

func NewTasksHandler(service *TasksService) *chi.Mux {
	handler := Handler{
		Service: service,
	}
	r := chi.NewRouter()
	r.Post("/create", handler.CreateTaskHandler)
	r.Post("/get", handler.GetAllTasksHandler)
	r.Get("/{id}", handler.GetTaskWithNoteHandler)
	r.Patch("/{id}/update", handler.ChangeTaskHandler)
	r.Delete("/{id}/delete", handler.DeleteTaskHandler)
	r.Post("/{id}/status/update", handler.UpdateStatusHandler)
	return r
}

// Criar uma task do começo
func (h Handler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {

	bodyRequest := &models.TaskDto{}
	resposta := map[string]any{"Sucess": true}
	ctx := context.Background()
	err := httphandler.ReadBody(r.Body, bodyRequest)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	id, err := h.Service.CreateTask(ctx, bodyRequest.Title, bodyRequest.Description, bodyRequest.Status, bodyRequest.DueDate, bodyRequest.Importante)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
		return
	}
	resposta["id"] = id
	httphandler.RespondSucess(resposta, w)
}

// Buscar todas as tasks com filtro
func (h *Handler) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {

	//Ler requisição
	bodyRequest := &models.TaskDto{}
	ctx := context.Background()
	resposta := map[string]any{"Sucess": true}
	err := httphandler.ReadBody(r.Body, bodyRequest)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	tasks, err := h.Service.GetAll(ctx, bodyRequest.Filter)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
		return
	}

	//Responder
	resposta["Tasks"] = tasks
	httphandler.RespondSucess(resposta, w)
}

// Alterar task pelo ID da task
func (h Handler) ChangeTaskHandler(w http.ResponseWriter, r *http.Request) {

	bodyRequest := &models.TaskDto{}
	resposta := map[string]any{"Sucess": true}
	ctx := context.Background()
	err := httphandler.ReadBody(r.Body, bodyRequest)
	id := chi.URLParam(r, "id")
	fmt.Println(id)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	resp, err := h.Service.ChangeTask(ctx, id, bodyRequest.Title, bodyRequest.Description, bodyRequest.DueDate, bodyRequest.Importante)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
	}
	fmt.Println(resp)
	//Responder
	resposta["Task"] = resp
	httphandler.RespondSucess(resposta, w)

}

// Deletar tasks pelo ID
func (h Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	//Ler Body
	bodyRequest := &models.TaskDto{}
	id := chi.URLParam(r, "id")
	fmt.Println(id)
	resposta := map[string]any{"Sucess": true}
	ctx := context.Background()
	err := httphandler.ReadBody(r.Body, bodyRequest)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	resp, err := h.Service.DeleteTask(ctx, id)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
	}
	//Responder
	resposta["Delete"] = resp
	httphandler.RespondSucess(resposta, w)
}

// Mudar status da task pelo ID
func (h Handler) UpdateStatusHandler(w http.ResponseWriter, r *http.Request) {

	bodyRequest := &models.TaskDto{}
	resposta := map[string]any{"Sucess": true}
	ctx := context.Background()
	err := httphandler.ReadBody(r.Body, bodyRequest)
	id := chi.URLParam(r, "id")
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	resp, err := h.Service.ChangeStatus(ctx, id, bodyRequest.Status)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
	}
	//Responder
	fmt.Println(resp)
	resposta["Task"] = resp
	httphandler.RespondSucess(resposta, w)

}

// Buscar a task pelo ID e buscar as notas que ela contém
func (h Handler) GetTaskWithNoteHandler(w http.ResponseWriter, r *http.Request) {

	bodyRequest := &models.TaskDto{}
	resposta := map[string]any{"Sucess": true}
	id := chi.URLParam(r, "id")
	ctx := context.Background()
	err := httphandler.ReadBody(r.Body, bodyRequest)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	resp, err := h.Service.GetTaskWithNote(ctx, id)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
		return
	}
	//Responder
	resposta["Task"] = resp
	httphandler.RespondSucess(resposta, w)

}
