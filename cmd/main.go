package main

import (
	"github.com/rafaelsouzaribeiro/rate-limiter-in-golang/configs"
	"github.com/rafaelsouzaribeiro/rate-limiter-in-golang/internal/infra/di"
)

func main() {

	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	di.NewDI()

}
