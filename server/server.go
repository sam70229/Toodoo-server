package server

import (
	"fmt"
	"net"
	"net/http"

	"Toodoo/database"
	"Toodoo/routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/knadh/koanf"
	"go.uber.org/zap"
)

type Server struct {
    logger *zap.SugaredLogger
    router chi.Router
    server *http.Server
}

func New(config *koanf.Koanf) (*Server, error) {
    r := chi.NewRouter()
    r.Use(middleware.RequestID)
    r.Use(middleware.Recoverer)
    
    r.Use(cors.New(cors.Options{
        AllowedOrigins:   config.Strings("server.cors.allowed_origins"),
		AllowedMethods:   config.Strings("server.cors.allowed_methods"),
		AllowedHeaders:   config.Strings("server.cors.allowed_headers"),
		AllowCredentials: config.Bool("server.cors.allowed_credentials"),
		MaxAge:           config.Int("server.cors.max_age"),
    }).Handler)

    s := &Server{
        logger: zap.S().With("package", "server"),
        router: r,
    }

    return s, nil
}

func (s *Server) Run(config *koanf.Koanf) error {
    s.logger.Info("Starting Server...")
    client, err := database.Connect()

    routes.NewRouter(s.Router(), client)

    s.server = &http.Server{
        Addr: net.JoinHostPort(config.String("server.host"), config.String("server.port")),
        Handler: s.router,
    }
    
    // listener, err := net.Listen("tcp", s.server.Addr)
    if err != nil {
        return fmt.Errorf("Could not listen on %s: %v", s.server.Addr, err)
    }

    // go s.server.Serve(listener)
    go s.server.ListenAndServe()
    s.logger.Infow("API Listening", "address", s.server.Addr)

    return nil
}

func (s *Server) Router() chi.Router {
    return s.router
}