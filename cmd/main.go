package main

import (
	"github.com/gookit/config/v2"

	"github.com/lapacek/simple-api-example/internal"
)

// Configuration related constants
const (
	confName = "server-config"
	confKeyToLower = false
)

// Represents vector of the postgres database connection related environment variables names
var pgEnv = []string{
	internal.PGUser,
	internal.PGPass,
	internal.PGName,
	internal.PGHost,
	internal.PGPort,
}

func main() {

	conf := config.NewEmpty(confName)
	conf.LoadOSEnv(pgEnv, confKeyToLower)

	server := internal.NewServer(conf)
	server.Run()
}


