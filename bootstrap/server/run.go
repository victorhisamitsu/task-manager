package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Hitsa/task-manager/bootstrap/database"
	"github.com/Hitsa/task-manager/internal/app/note"
	"github.com/Hitsa/task-manager/internal/app/tasks"
)

func Run() {
	// Conectar com o DB

	db := database.ConexaoDb()
	database.CreateTables(context.Background(), db)

	// Fornecer o acesso ao banco de dados para a Repository
	NoteRepository := note.NewRepositoryNote(db)

	TasksRepository := tasks.NewRepositoryTasks(db)

	// Forcener a repository para a service
	noteService := note.NewNoteService(NoteRepository)

	taskService := tasks.NewTasksService(TasksRepository, noteService)

	// Fornecer a service para a handler
	router := SetupHttpServer(taskService, noteService)
	fmt.Println(router)
	http.ListenAndServe(":8080", router)
}
