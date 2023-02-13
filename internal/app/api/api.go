package api

import (
	"github.com/julienschmidt/httprouter"
	_cors "github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	config *Config
	logger *logrus.Logger
	router *httprouter.Router
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: httprouter.New(),
	}
}

func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	cors := _cors.New(_cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	router := cors.Handler(s.router)
	s.logger.Info("Start server")

	return http.ListenAndServe(s.config.BindAddr, router)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *Server) configureRouter() {
	s.router.HandlerFunc("GET", "/healthchecker", s.logMiddleware(s.handleHealthChecker()))
	s.router.HandlerFunc("GET", "/send-analytics", s.logMiddleware(s.handleSendAnalytics()))
	s.router.HandlerFunc("GET", "/get-analytics", s.logMiddleware(s.handleReadAll()))
}
