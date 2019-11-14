package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"os"
	"os/signal"
	"syscall"
)


//
func main() {
	r := gin.Default()

	r.GET("/rest", func(c *gin.Context) {

	})

	srv := &http.Server{
		Addr: ":8585",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shut")
}
