package middleware

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func JSONType(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Header().Set("Content-Type", "application/repos.v1+json; charset=utf-8")
		next(w, r, params)
	}
}
