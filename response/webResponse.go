package response

type Response struct {
	Code   int
	Status string
	Data   interface{} `json:"data,omitempty`
}

func RespData(code int, status string, data interface{}) Response {
	return Response{
		Code:   code,
		Status: status,
		Data:   data,
	}
}
