package handler

import (
	"encoding/json"
	"go-docs/internal/shared"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		result := shared.ApiResponse{
			Status: "error",
			Code:   "METHOD_NOT_ALLOWED",
			Messages: []shared.Message{
				{
					Lang:    "en",
					Message: "Only GET method is allowed",
				},
			},
		}
		json.NewEncoder(w).Encode(result)
		return
	}

	result := shared.ApiResponse{
		Status: "success",
		Code:   "SERVER_LIFE",
		Messages: []shared.Message{
			{
				Lang:    "en",
				Message: "Success request",
			},
		},
	}
	json.NewEncoder(w).Encode(result)
}
