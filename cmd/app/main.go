package main

import (
	"fmt"
	"log"

	"CapitalParser/internal/config"
)

func main() {
	config.LoadConfig()
	fmt.Println("Starting Capital.com Market Monitor...")

	// TODO: Initialize configuration
	// TODO: Set up logger
	// TODO: Create API client
	// TODO: Initialize storage
	// TODO: Setup analyzer
	// TODO: Start data collector

	log.Println("Application running. Press Ctrl+C to stop.")

	// TODO: Add graceful shutdown
	select {}
}
