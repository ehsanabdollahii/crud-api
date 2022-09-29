package main

import (
	"crud-api/config"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
	"time"
)

// setupLogging sets logging level for logrus
func setupLogging() {
	switch strings.ToLower(config.Logging.Level) {
	case "debug":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}

// cliFlags returns global cli flags
func cliFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "debug",
			Value:       false,
			Usage:       "Activate debug information",
			EnvVars:     []string{"DEBUG"},
			Destination: &config.Server.Debug,
		},
		&cli.StringFlag{
			Name:        "logging-level",
			Value:       "info",
			Usage:       "set logging level",
			EnvVars:     []string{"LOG_LEVEL"},
			Destination: &config.Logging.Level,
		},
	}
}

func main() {
	app := &cli.App{
		Name:     "ehsan",
		Usage:    "ehsan",
		Compiled: time.Now(),
		Version:  "0.1",
		Authors: []*cli.Author{
			{
				Name:  "Ehsan Abdollahi",
				Email: "ehsaaan.a@gmail.com",
			},
		},
		Flags: cliFlags(),
		Commands: []*cli.Command{
			Server(),
		},
		Before: func(c *cli.Context) error {
			setupLogging()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal("failed starting the web server")
	}
}
