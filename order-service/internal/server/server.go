package server

// import (
// 	"log"

// 	"github.com/gin-gonic/gin"
// )

// type Server interface {
// 	Start() error
// }

// type server struct {
// 	router *gin.Engine
// 	cfg    *Config
// }

// func NewServer(cfg *Config) Server {
// 	r := gin.Default()

// 	return &server{
// 		router: r,
// 		cfg:    cfg,
// 	}
// }

// func (s *server) Start() error {
// 	if err := s.registerRoutes(); err != nil {
// 		log.Printf("Error registering routes: %v", err)
// 		return err
// 	}

// 	log.Printf("Starting server on port %s...", s.cfg.Port)
// 	if err := s.router.Run(s.cfg.Port); err != nil {
// 		log.Fatalf("Failed to start server: %v", err)
// 		return err
// 	}

// 	return nil
// }
