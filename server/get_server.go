package server

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
)


func SetServer(app *echo.Echo) {
	server := &http2.Server{
        MaxConcurrentStreams: 250,
        MaxReadFrameSize:     1048576,
        IdleTimeout:          10 * time.Second,
    }
    err := app.StartH2CServer(":3000", server)
    if  err != http.ErrServerClosed {
        log.Fatal(err)
      }
}


