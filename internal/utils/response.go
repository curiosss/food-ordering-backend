package utils

type Response struct {
	Status  bool   `json:"status"`
	Errors  string `json:"errors"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
