package http_service

var NotFoundResponse = NewNotFoundResponse("Endpoint not found")

type ErrorFieldResponseList []ErrorFieldResponse

func (list *ErrorFieldResponseList) Add(errorFieldResponseList ...ErrorFieldResponse) {
	for _, errorFieldResponse := range errorFieldResponseList {
		*list = append(*list, errorFieldResponse)
	}
}

func (list *ErrorFieldResponseList) Contains(errorField string) bool {
	for _, fieldResponse := range *list {
		if fieldResponse.Field == errorField {
			return true
		}
	}

	return false
}

type ErrorFieldResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field"`
}

type ErrorResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Errors  []ErrorFieldResponse `json:"errors,omitempty"`
}

func NewErrorResponse(code int, message string, errors []ErrorFieldResponse) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: message,
		Errors:  errors,
	}
}

func NewNotFoundResponse(message string) ErrorResponse {
	return ErrorResponse{
		Code:    40400,
		Message: message,
		Errors:  nil,
	}
}

func NewForbiddenResponse(message string) ErrorResponse {
	return ErrorResponse{
		Code:    40300,
		Message: message,
		Errors:  nil,
	}
}

func NewUnauthorizedResponse(message string) ErrorResponse {
	return ErrorResponse{
		Code:    40100,
		Message: message,
		Errors:  nil,
	}
}

func NewValidationErrorResponse(errors []ErrorFieldResponse) ErrorResponse {
	return ErrorResponse{
		Code:    40000,
		Message: "Validation errors",
		Errors:  errors,
	}
}

func NewErrorFieldResponse(code int, message string, field string) ErrorFieldResponse {
	return ErrorFieldResponse{
		Code:    code,
		Message: message,
		Field:   field,
	}
}
