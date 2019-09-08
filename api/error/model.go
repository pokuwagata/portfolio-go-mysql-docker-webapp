package error

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
}