package handler

import (
	"fmt"
	"net/http"

	"github.com/h-tachikawa/mechanical-receptionist/usecase"
)

type NotificationHandler struct{}

func NewNotificationHandler() NotificationHandler {
	return NotificationHandler{}
}

func ensurePostRequest(r *http.Request) bool {
	if r.Method == http.MethodPost {
		return true
	}
	return false
}

func (n NotificationHandler) Handle(w http.ResponseWriter, r *http.Request) {
	isPostRequest := ensurePostRequest(r)

	if !isPostRequest {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Only allowed POST method.")
	}

	err := usecase.NewNotificationUseCase().Execute()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "error")
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}
