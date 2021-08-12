package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ezegrosfeld/vader"
	"github.com/ezegrosfeld/yoda"
)

type response struct {
	Data interface{} `json:"data"`
}

func main() {
	server := yoda.NewServer(yoda.Config{
		Addr:         ":8080",
		Name:         "M1",
		IdleTimeout:  time.Second * 30,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	})

	r := server.Group("/ping")
	r.Get("", func(c *yoda.Context) error {
		// Make request to m2 server
		res, err := http.Get(fmt.Sprintf("%s/ping", os.Getenv("M2_URL")))
		/* res, err := http.Get("http://m2/ping") */
		if err != nil {
			return c.Error(vader.Internal(err.Error()))
		}

		// Parse res body to a response struct using json package
		var response response
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			return c.Error(vader.Internal(err.Error()))
		}

		// Return response struct as json
		return c.JSON(http.StatusOK, response.Data)
	})

	server.Start()
}
