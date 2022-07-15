package utils

//Response Struct
type Response struct {
	Result  bool        `json:"result"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// Response Model
func HandleResponseModel(result bool, message string, errors interface{}, data interface{}) Response {
	return Response{
		Result:  result,
		Message: message,
		Errors:  errors,
		Data:    data,
	}
}
