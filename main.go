package main

import (
	"flag"
	"os"
)

const defaultHandler = "custom"

func main() {
	port, customHandler := parseFlags()

	server := Server{port: port, customHandler: customHandler}

	if err := server.Start(); err != nil {
		panic(err)
	}
}

func parseFlags() (int, string) {
	port := flag.Int("port", 8080, "port to start redirect server on")
	customHandlerCLI := flag.String("handler", defaultHandler, "custom handler to redirect to")
	flag.Parse()
	customHandlerEnv := os.Getenv("REDIRECT_HANDLER")
	if customHandlerEnv == "" {
		return *port, *customHandlerCLI
	}
	if *customHandlerCLI != defaultHandler {
		return *port, *customHandlerCLI
	}
	return *port, customHandlerEnv
}
