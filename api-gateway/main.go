package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API Gateway is healthy"})
	})

	inventoryServiceURL := "http://inventory-service:8010"
	router.Any("/inventory/*path", reverseProxy(inventoryServiceURL))

	orderServiceURL := "http://order-service:8020"
	router.Any("/orders/*path", reverseProxy(orderServiceURL))

	log.Println("Starting API Gateway on port 8000...")
	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Failed to start API Gateway: %v", err)
	}
}

func reverseProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetURL, err := url.Parse(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid target URL"})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
			log.Printf("Proxy error: %v", err)
			rw.WriteHeader(http.StatusBadGateway)
			rw.Write([]byte("Bad Gateway"))
		}

		c.Request.URL.Path = c.Param("path")
		c.Request.Host = targetURL.Host
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
