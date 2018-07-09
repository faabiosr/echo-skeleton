package main

import (
	"github.com/fabiorphp/echo-skeleton/pkg/cli"
	ufcli "github.com/urfave/cli"
	"os"
)

var (
	appName = "api"
	version = "0.0.0"
)

func main() {
	app := ufcli.NewApp()
	app.Name = appName
	app.Version = version
	app.Usage = "Manage something"
	app.Flags = []ufcli.Flag{
		ufcli.StringFlag{
			Name:   "listen, l",
			Value:  "0.0.0.0:8000",
			Usage:  "Address and port on which App will accept HTTP requests",
			EnvVar: "LISTEN",
		},
		ufcli.StringFlag{
			Name:   "log-folder, lf",
			Value:  "",
			Usage:  `Log folder path for access and application logging. Default "stdout"`,
			EnvVar: "LOG_FOLDER",
		},
	}

	app.Action = cli.Serve

	_ = app.Run(os.Args)
}
