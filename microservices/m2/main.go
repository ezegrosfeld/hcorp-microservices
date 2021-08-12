package main

import (
	"time"

	"github.com/ezegrosfeld/yoda"
)

func main() {
	server := yoda.NewServer(yoda.Config{
		Addr:         ":8080",
		Name:         "M2",
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	})

	r := server.Group("/ping")
	r.Get("", func(c *yoda.Context) error {
		return c.JSON(200, "pong")
	})

	server.Start()
}
