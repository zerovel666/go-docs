package shared

type ApiResponse struct {
	Status   string    `json:"status"`
	Code     string    `json:"code"`
	Messages []Message `json:"messages"`
	Data     any       `json:"data"`
}

type Message struct {
	Lang    string `json:"lang"`
	Message string `json:"message"`
}
