package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	_, err := ctx.WriteString("Goafka!")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func main() {
	port := "8081"
	err := fasthttp.ListenAndServe(":"+port, fastHTTPHandler)
	log.Printf("Server started on port: %s", port)
	if err != nil {
		log.Fatalf("ERROR: starting server %s", err)
	}
}
