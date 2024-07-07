package httphandlers

import (
	"encoding/json"
	"net/http"
)

type HttpTemplate struct {
}

func NewHttpTemplate() *HttpTemplate {
	return &HttpTemplate{}
}

func (h *HttpTemplate) Write(w http.ResponseWriter, _return interface{}, httpStatus int) {
	w.Header().Set("Content-Type", "application/json")

	b, err := json.Marshal(_return)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(httpStatus)
	_, _ = w.Write(b)
}

func (h *HttpTemplate) WriteError(w http.ResponseWriter, httpErr *Error) {
	w.Header().Set("Content-Type", "application/json")

	b, err := json.Marshal(httpErr)
	if err != nil {
		w.WriteHeader(httpErr.Status)
		b, _ := json.Marshal(ErrUnknown)
		_, _ = w.Write(b)
		return
	}

	w.WriteHeader(httpErr.Status)
	_, _ = w.Write(b)
}
