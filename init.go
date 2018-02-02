package main

import (
	"flag"

	"github.com/olebedev/config"
)

var (
	cfg *config.Config

	flagWebhook = flag.Bool("webhook", false, "enable webhook mode")
)

func init() {
	var err error
	cfg, err = config.ParseYamlFile("./config.yaml")
	errCheck(err)
}
