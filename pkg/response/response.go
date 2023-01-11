//paquete de respuestas 

package response

type Response struct{
	message string
	data    interface{}
}

func OK(message string, data interface{}) *Response {
	return &Response{
        message: message,
        data:    data,
    }
}

func Err(err error) *Response {
	return &Response{
        message: err.Error(),
		data:    nil,
    }
}