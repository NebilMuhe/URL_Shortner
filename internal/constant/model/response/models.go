package response

type Response struct {
	Ok    bool           `json:"ok"`
	Data  any            `json:"data,omitempty"`
	Error *ErrorResponse `json:"error,omitempty"`
}

type ErrorResponse struct {
	StausCode  int          `json:"status_code"`
	Message    string       `json:"message"`
	FieldError []FieldError `json:"field_error,omitempty"`
}

type FieldError struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
