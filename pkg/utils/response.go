package utils

//Response struct for json response
type Response struct {
	Result  bool        `json:"result"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// for response model
func HandleResponseModel(result bool, message string, errors interface{}, data interface{}) Response {
	return Response{
		Result:  result,
		Message: message,
		Errors:  errors,
		Data:    data,
	}
}
