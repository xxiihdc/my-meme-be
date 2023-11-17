package response

type Response struct {
	Code   int
	Status string
	Data   interface{} `json:"data,omitempty`
}
