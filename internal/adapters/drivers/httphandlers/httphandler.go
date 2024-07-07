package httphandlers

import "net/http"

type HttpHandler interface {
	Method() string
	Pattern() string
	Func(w http.ResponseWriter, r *http.Request)
}
