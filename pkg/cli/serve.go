package cli

import (
	"context"
	"github.com/coreos/go-systemd/activation"
	"github.com/fabiorphp/echo-skeleton/pkg/handler"
	lg "github.com/fabiorphp/echo-skeleton/pkg/log"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"time"
)

// Serve Http Server
func Serve(c *cli.Context) error {

	// Echo instance
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.HTTPErrorHandler = handler.Error
	e.Logger.SetLevel(log.INFO)
	e.Logger.SetOutput(lg.File(c.String("log-folder") + "/app.log"))

	logConfig := mw.DefaultLoggerConfig
	logConfig.Output = lg.File(c.String("log-folder") + "/access.log")

	e.Pre(mw.RemoveTrailingSlash())
	e.Use(mw.LoggerWithConfig(logConfig))
	e.Use(mw.Recover())
	e.Use(mw.RequestID())
	e.Use(mw.Secure())

	// Index
	e.GET("/", handler.Index())

	// Start server
	e.Logger.Infof("%s service - %s", c.App.Name, c.App.Version)

	go func() {
		if err := start(e, c); err != nil {
			e.Logger.Info("shutting down the server")
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

	return nil
}

// Starts http server when SystemD socket is defined
func start(e *echo.Echo, c *cli.Context) error {
	listeners, err := activation.Listeners()

	if err != nil {
		return err
	}

	address := c.String("listen")

	if len(listeners) > 0 {
		e.Listener = listeners[0]

		address = e.Listener.Addr().String()
	}

	e.Logger.Infof("starting server on %s", address)

	return e.Start(c.String("listen"))
}
