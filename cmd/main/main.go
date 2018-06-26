package main

import (
	"log"
	"strings"

	"github.com/anatollupacescu/auth-test/internal/app/tool"
	"github.com/spf13/viper"
)

const Endpoints = "ENDPOINT_LIST"

func main() {
	viper.AutomaticEnv()

	endpoints := viper.GetString(Endpoints)

	if len(endpoints) < 1 {
		log.Fatalf("endpoint list not found in environment variable %s, exiting...", Endpoints)
		return
	}

	list := strings.Split(endpoints, ",")

	for _, url := range list {
		doTest(url)
	}
}

func doTest(url string) {
	endpoint, err := tool.NewEndpoint(url)
	if err != nil {
		log.Printf("could not create endpoint: %s", err)
	}
	responseCode, err := endpoint.ResponseCode()
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("request to url '%s' returned code: %d", url, responseCode)
	}
}
