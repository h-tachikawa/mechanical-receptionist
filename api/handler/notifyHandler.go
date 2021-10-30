package handler

import (
	"fmt"
	"net/http"

	"github.com/h-tachikawa/mechanical-receptionist/api/usecase"
)

func NotifyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Only allowed POST method.")
		return
	}

	err := usecase.NotifyUseCase()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "error")
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}
