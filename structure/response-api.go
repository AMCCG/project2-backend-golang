package structure

type ResponseApi struct {
	Status  string
	Code    int
	Message string
	Content interface{}
}
