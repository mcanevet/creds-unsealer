package config

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/jessevdk/go-flags"
)

// Config stores the handler's configuration and UI interface parameters
type Config struct {
	Version   bool     `short:"V" long:"version" description:"Display version."`
	LogLevel  string   `short:"l" long:"loglevel" description:"Set loglevel ('debug', 'info', 'warn', 'error', 'fatal', 'panic')." env:"BIVAC_LOG_LEVEL" default:"info"`
	Manpage   bool     `short:"m" long:"manpage" description:"Output manpage."`
	Backend   string   `short:"b" long:"backend" description:"Backend to use." env:"CREDS_BACKEND" default:"pass"`
	Providers []string `short:"p" long:"providers" description:"Providers to use." env:"CREDS_PROVIDERS" default:"ovh"`
	Pass      struct {
		Path string `long:"pass-path" description:"Path to password-store." env:"CREDS_PASS_PATH"`
	} `group:"Pass backend options"`

	Provider struct {
		InputPath string `long:"provider-input-path" description:"Provider input path"`
	} `group:"Provider options"`
}

// LoadConfig loads the config from flags & environment
func LoadConfig(version string) *Config {
	var c Config
	parser := flags.NewParser(&c, flags.Default)
	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}

	if c.Version {
		fmt.Printf("Creds-unsealer %v\n", version)
		os.Exit(0)
	}

	if c.Manpage {
		os.Exit(0)
	}

	err := c.setupLogLevel()
	if err != nil {
		log.Errorf("failed to setup log level: %s", err)
		os.Exit(1)
	}
	return &c
}

func (c *Config) setupLogLevel() (err error) {
	switch c.LogLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		err = fmt.Errorf("Wrong log level '%v'", c.LogLevel)
	}

	return
}
