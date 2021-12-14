package response

import (
	"net/http"
)

// function response bad request
func BadRequestResponse(message string) map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusBadRequest,
		"Message": message,
	}
	return result
}

// function response success dengan return value data
func SuccessResponseData(message string, data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusOK,
		"Message": message,
		"Data":    data,
	}
	return result
}

// function response success dengan return value data
func SuccessResponseDataOrder(message string, data interface{}, user interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusOK,
		"Message": message,
		"Data":    data,
		"User":    user,
	}
	return result
}

// function response success tanpa return value data
func SuccessResponseNonData(message string) map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusOK,
		"Message": message,
	}
	return result
}
