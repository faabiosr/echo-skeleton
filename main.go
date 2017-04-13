package main

import (
	"context"
	"github.com/fabiorphp/myapp/handler"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/color"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"time"
)

var (
	version string = "0.0.0"
	appName string = "myapp"
	port    string
)

const (
	MYAPP_PORT = "MYAPP_PORT"
)

func init() {
	port = os.Getenv(MYAPP_PORT)

	if port == "" {
		port = "8888"
	}
}

func main() {

	// Echo instance
	e := echo.New()
	e.HTTPErrorHandler = handler.Error
	e.Logger.SetLevel(log.INFO)

	// Middlewares
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Secure())

	// Routes => handler
	e.GET("/", handler.HomeIndex)

	// Start server
	colorer := color.New()
	colorer.Printf("⇛ %s service - %s\n", appName, color.Green(version))

	go func() {
		if err := e.Start(":" + port); err != nil {
			colorer.Printf(color.Red("⇛ shutting down the server\n"))
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
