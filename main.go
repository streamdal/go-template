package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/batchcorp/go-template/api"
	"github.com/batchcorp/go-template/config"
	"github.com/batchcorp/go-template/deps"
)

var (
	version = "No version specified"

	envFile = kingpin.Flag("envfile", "Local Env file to read at startup").Short('e').Default(".env").String()
	debug   = kingpin.Flag("debug", "Enable debug output").Short('d').Bool()
)

func init() {
	// Parse CLI stuff
	kingpin.Version(version)
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.CommandLine.VersionFlag.Short('v')
	kingpin.Parse()
}

func main() {
	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// JSON formatter for log output if not running in a TTY
	// because Loggly likes JSON but humans like colors
	if !terminal.IsTerminal(int(os.Stderr.Fd())) {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	log := logrus.WithField("method", "main")
	log.WithField("filename", *envFile).Debug("Loading env file")

	if err := godotenv.Load(*envFile); err != nil {
		log.WithFields(logrus.Fields{"filename": *envFile, "err": err.Error()}).Warn("Unable to load dotenv file")
	}

	cfg := config.New()
	if err := cfg.LoadEnvVars(); err != nil {
		log.WithError(err).Fatal("Could not instantiate configuration")
	}

	d, err := deps.New(cfg)
	if err != nil {
		log.WithError(err).Fatal("Could not setup dependencies")
	}

	go d.RabbitMQ.Listen()

	log = log.WithField("environment", cfg.EnvName)

	a := api.New(cfg, d, version)
	log.Fatal(a.Run())

}
