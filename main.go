package main

import (
	"net/http"
	"log"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"
)

func main() {
	addr := ":8080"
	router := httprouter.New()
	svc := stringService{}

	countHandler := routerWrapper(
		httpTransport.NewServer(
			makeCountEndpoint(svc),
			decodeHttpRequest,
			encodeHttpResponse,
		),
	)

	lowerCaseHandler := routerWrapper(
		httpTransport.NewServer(
			makeUpperCaseEndpoint(svc),
			decodeHttpRequest,
			encodeHttpResponse,
		),
	)

	upperCaseHandler := routerWrapper(
		httpTransport.NewServer(
			makeUpperCaseEndpoint(svc),
			decodeHttpRequest,
			encodeHttpResponse,
		),
	)

	router.POST("/count", countHandler)
	router.POST("/lowercase", lowerCaseHandler)
	router.POST("/uppercase", upperCaseHandler)

	log.Printf("Listeing on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, router))
}
