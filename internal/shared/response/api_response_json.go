package response

import (
	"encoding/json"
	"go-docs/internal/shared/dto"
	"net/http"
)

func ApiResponseJSON(
	w http.ResponseWriter,
	r *http.Request,
	code string,
	messageRU string,
	messageEN string,
	data interface{},
) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		dto.ApiResponse{
			Status: "success",
			Code:   code,
			Messages: []dto.Message{
				{
					Lang:    "ru",
					Message: messageRU,
				},
				{
					Lang:    "en",
					Message: messageEN,
				},
			},
			Data: data,
		},
	)
}
