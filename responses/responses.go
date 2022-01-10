package responses

import "net/http"

// function for success response non data
func SuccessResponseNonData(message string) map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusOK,
		"Message": message,
	}
	return result
}

// function for success response with return data
func SuccessResponseData(message string, data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusOK,
		"Message": message,
		"Data":    data,
	}
	return result
}

// function for bad request response
func BadRequestResponse(message string) map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusBadRequest,
		"Message": message,
	}
	return result
}

// function for internal server error response
func InternalServerError(message string) map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusInternalServerError,
		"Message": message,
	}
	return result
}

// function for data not found response
func StatusDataNotFound(message string) map[string]interface{} {
	result := map[string]interface{}{
		"Code":    http.StatusNotFound,
		"Message": message,
	}
	return result
}
