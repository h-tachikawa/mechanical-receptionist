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

func (n NotificationHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Only allowed POST method.")
		return
	}

	err := usecase.NewNotificationUseCase().Execute()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "error")
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}
