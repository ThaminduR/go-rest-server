package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Config structures
type ServerConfig struct {
	Port string `json:"port"`
}

type EndpointResponse struct {
	Status int                    `json:"status"`
	Body   map[string]interface{} `json:"body"`
}

type Endpoint struct {
	Path     string           `json:"path"`
	Method   string           `json:"method"`
	Response EndpointResponse `json:"response"`
}

type Config struct {
	Server    ServerConfig `json:"server"`
	Endpoints []Endpoint   `json:"endpoints"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

var config Config

// Load configuration from JSON file
func loadConfig(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return fmt.Errorf("failed to parse config file: %v", err)
	}

	return nil
}

// Create a handler for configured endpoints
func createConfiguredHandler(endpoint Endpoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if method matches
		if r.Method != endpoint.Method {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Method not allowed"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(endpoint.Response.Status)
		
		// Return only what's defined in config
		json.NewEncoder(w).Encode(endpoint.Response.Body)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(ErrorResponse{Error: "Not Found"})
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s", r.Method, r.URL.Path)
		next(w, r)
		log.Printf("Request completed in %v", time.Since(start))
	}
}

func main() {
	// Load configuration
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		configFile = "config.json"
	}

	log.Printf("Loading configuration from: %s", configFile)
	err := loadConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Get port from environment or config
	port := os.Getenv("PORT")
	if port == "" {
		port = config.Server.Port
	}
	if port == "" {
		port = "8080"
	}

	// Register default 404 handler
	http.HandleFunc("/", loggingMiddleware(notFoundHandler))

	// Register all configured endpoints
	log.Printf("Configuring endpoints from config file:")
	for _, endpoint := range config.Endpoints {
		handler := createConfiguredHandler(endpoint)
		http.HandleFunc(endpoint.Path, loggingMiddleware(handler))
		log.Printf("  %s %s -> Status %d", endpoint.Method, endpoint.Path, endpoint.Response.Status)
	}

	addr := ":" + port
	log.Printf("\nServer starting on http://localhost%s", addr)
	log.Printf("Total endpoints configured: %d", len(config.Endpoints))
	log.Printf("Press Ctrl+C to stop\n")
	
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
