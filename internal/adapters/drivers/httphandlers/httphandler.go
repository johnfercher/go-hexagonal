package httphandlers

import "net/http"

type HTTPHandler interface {
	Method() string
	Pattern() string
	Func(w http.ResponseWriter, r *http.Request)
}
