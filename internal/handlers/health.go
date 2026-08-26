package handler

import (
	"go-docs/internal/shared/response"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	response.ApiResponseJSON(
		w,
		r,
		"SERVER_LIFE",
		"Сервер в рабочем состоянии",
		"The server is in work condition",
		nil,
	)
}
