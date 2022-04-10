package main

import (
	"github.com/gookit/config/v2"
	"github.com/lapacek/simple-api-example/internal"
	"github.com/lapacek/simple-api-example/internal/common"
)

// Configuration related constants
const (
	confName = "server-config"
	confKeyToLower = false
)

// Represents vector of the postgres database connection related environment variables names
var pgEnv = []string{
	common.PGUser,
	common.PGPass,
	common.PGName,
	common.PGHost,
	common.PGPort,
}

func main() {

	conf := config.NewEmpty(confName)
	conf.LoadOSEnv(pgEnv, confKeyToLower)

	server := internal.NewServer(conf)
	server.Run()
}


