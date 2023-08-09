package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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

func (h Handler) BodyApiHandlerTasks(w http.ResponseWriter, r *http.Request) {
	resposta := map[string]any{}
	body, _ := io.ReadAll(r.Body)
	var minhaVariavel map[string]string
	json.Unmarshal(body, &minhaVariavel)
	resposta["resposta"] = minhaVariavel
	variavelJson, _ := json.Marshal(minhaVariavel)
	w.Write(variavelJson)
}

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

func (h Handler) ChangeTaskHandler(w http.ResponseWriter, r *http.Request) {

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
	resp, err := h.Service.ChangeTask(ctx, id, bodyRequest.Title, bodyRequest.Description, bodyRequest.DueDate, bodyRequest.Importante)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
	}
	//Responder
	resposta["Task"] = resp
	httphandler.RespondSucess(resposta, w)

}

func (h Handler) GetTaskHandler(w http.ResponseWriter, r *http.Request) {

	bodyRequest := &models.TaskDto{}
	resposta := map[string]any{"Sucess": true}
	ctx := context.Background()
	err := httphandler.ReadBody(r.Body, bodyRequest)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	resp, err := h.Service.GetTask(ctx, bodyRequest.ID)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
	}
	//Responder
	resposta["Task"] = resp
	httphandler.RespondSucess(resposta, w)

}

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
	resposta["Task"] = resp
	httphandler.RespondSucess(resposta, w)
}

// TODO
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
	resp, err := h.Service.ChangeTask(ctx, id, bodyRequest.Title, bodyRequest.Description, bodyRequest.DueDate, bodyRequest.Importante)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
	}
	//Responder
	resposta["Task"] = resp
	httphandler.RespondSucess(resposta, w)

}

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
	resp, err := h.Service.GetTask(ctx, id)
	if err != nil {
		httphandler.RespondError(err.Error(), resposta, w)
	}
	//Responder
	resposta["Task"] = resp
	httphandler.RespondSucess(resposta, w)

}
