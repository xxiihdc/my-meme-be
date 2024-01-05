package orm

type Model interface {
	FromMap(data map[string]interface{}) error
	JSONToModel(jsonData []byte, model Model) error
}
