package server

import (
	"fmt"
	"log"

	cache "inventory-service/internal/adapter/cahce"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Start() error
}

type server struct {
	router *gin.Engine
	cfg    *Config
	cache  cache.Cache
}

func NewServer(cfg *Config) (Server, error) {
	r := gin.Default()

	// Initialize Redis cache
	redisCache, err := cache.NewRedisCache(cfg.RedisAddr, cfg.RedisPass, cfg.RedisDB)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize redis cache: %w", err)
	}

	return &server{
		router: r,
		cfg:    cfg,
		cache:  redisCache,
	}, nil
}

func (s *server) Start() error {
	if err := s.registerRoutes(); err != nil {
		log.Printf("Error registering routes: %v", err)
		return err
	}

	log.Printf("Starting server on port %s...", s.cfg.Port)
	if err := s.router.Run(s.cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return err
	}

	return nil
}
