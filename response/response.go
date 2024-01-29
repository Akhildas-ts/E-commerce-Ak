package response

type Response struct {
	Statuscode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}

type SuccesResponse struct {
	Statuscode int         `json:"status_code"`
	Message    string      `json:"message"`
}

func ClientResponse(statusCode int, message string, data interface{}, err interface{}) Response {
	return Response{
		Statuscode: statusCode,
		Message:    message,
		Data:       data,
		Error:      err,
	}

}

func SuccessClientResponse(statusCode int, message string) SuccesResponse{
	return SuccesResponse{
		Statuscode: statusCode,
		Message:    message,
	}
}
