package main

import (
	"flag"
	"os"
)

const defaultHandler = "custom"

func main() {
	customHandler := parseFlags()

	server := Server{customHandler: customHandler}

	if err := server.Start(); err != nil {
		panic(err)
	}
}

func parseFlags() string {
	customHandlerCLI := flag.String("handler", defaultHandler, "custom handler to redirect to")
	flag.Parse()
	customHandlerEnv := os.Getenv("REDIRECT_HANDLER")
	if customHandlerEnv == "" {
		return *customHandlerCLI
	}
	if *customHandlerCLI != defaultHandler {
		return *customHandlerCLI
	}
	return customHandlerEnv
}
