package http_server

import (
	"context"
	"net/http"
)

type ServerConfig struct {
	Port string
}

func NewServerConfig(port string) ServerConfig {
	return ServerConfig{
		Port: port,
	}
}

type Server struct {
	server http.Server
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown() error {
	return s.server.Shutdown(context.TODO())
}

func NewServer(conf ServerConfig, handler http.Handler) *Server {
	//fmt.Printf("Port : %s", fmt.Sprint(":"+conf.Port))
	serv := http.Server{
		Addr:    ":" + conf.Port,
		Handler: handler,
	}

	return &Server{server: serv}
}
