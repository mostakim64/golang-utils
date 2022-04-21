package logger

type SlacklogRequest struct {
	Message string `json:"message"`
	File    string `json:"file"`
	Level   string `json:"level"`
}
