module github.com/lapacek/simple-api-example/cmd

require (
	github.com/lapacek/simple-api-example/internal v0.0.0
)

replace github.com/lapacek/simple-api-example/internal => ../internal

go 1.16