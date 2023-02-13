package api

import (
	"log"
	"net/http"
)

func (s *Server) logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Print(request.Method, " -- ", request.URL)
		next.ServeHTTP(writer, request)
	}
}
