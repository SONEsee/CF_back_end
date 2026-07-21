package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/SONEsee/go-echo/api/routes"
	"github.com/SONEsee/go-echo/api/validators"
	bodylimit "github.com/SONEsee/go-echo/config/body-limit"
	"github.com/SONEsee/go-echo/config/cors"
	"github.com/SONEsee/go-echo/config/dotenv"
	"github.com/SONEsee/go-echo/config/loggers"
	recoverMiddleware "github.com/SONEsee/go-echo/config/recover"
	requestid "github.com/SONEsee/go-echo/config/request-id"
	"github.com/SONEsee/go-echo/config/secure"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	"github.com/labstack/echo/v4"
)

func init() {
	mode := os.Getenv("GO_ENV")
	if mode == "" {
		dotenv.LoadEnv()
		mode = os.Getenv("GO_ENV")
	}

	log.Println("Service run in mode", mode)
}

func main() {
	e := echo.New()

	//loggers
	e.Use(loggers.SetEchoLogger)

	//cors middilewares
	cors.SetCorsMiddlwares(e)

	//body limit
	bodylimit.SetBodyLimit(e)

	//request-id
	requestid.SetRequestID(e)

	//recover
	recoverMiddleware.SetRecoverMiddleware(e)

	//validator
	validators.Init()

	//secure
	secure.SetSecureMiddilware(e)

	e.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok sone!")
	})

	//serve uploaded files (images) statically
	e.Static("/uploads", "uploads")

	apiV1 := e.Group("/api/v1")
	routes.SetRoutes(apiV1)

	//connect to datbase
	dbpkg.CreateDatabaseConnection()

	//Error handler
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		status := http.StatusInternalServerError
		msg := err.Error()

		if he, ok := err.(*echo.HTTPError); ok {
			status = he.Code
			if m, ok := he.Message.(string); ok {
				msg = m
			}
		}
		// Send JSON response
		res := map[string]interface{}{
			"timestamp": time.Now().Format("2006-01-02-15-04-05"),
			"status":    status,
			"items":     nil,
			"error":     msg,
		}

		c.JSON(status, res)
	}

	//source: https://echo.labstack.com/docs/cookbook/graceful-shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		port := os.Getenv("PORT")
		if err := e.Start(fmt.Sprint(":", port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	log.Println("Running cleanup tasks...")
	log.Println("Echo was successful shutdown.")
}
