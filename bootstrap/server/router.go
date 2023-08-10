package server

import (
	"github.com/Hitsa/task-manager/internal/app/note"
	"github.com/Hitsa/task-manager/internal/app/tasks"
	"github.com/go-chi/chi/v5"
)

func SetupHttpServer(tasksService *tasks.TasksService, noteService *note.NoteService) *chi.Mux {
	router := chi.NewRouter()
	note := note.NewNotesHandler(noteService)
	tasks := tasks.NewTasksHandler(tasksService)

	router.Mount("/notes", note)
	router.Mount("/tasks", tasks)

	return router
}
