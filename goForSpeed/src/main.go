package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"sync"
	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for WebSocket connections
	},
}

// Middleware for CORS
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// HTTP Handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! Welcome to my HTTP server running on port %s.", flag.Lookup("port").Value.String())
}

func h1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handler 1.")
}

func h2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handler 2.")
}

func h3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handler 2.")
}

// WebSocket Handler
func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // Upgrade HTTP to WebSocket
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	fmt.Println("WebSocket connection established!")

	// Handle WebSocket messages
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("WebSocket error:", err)
			break
		}
		fmt.Printf("Received: %s\n", message)

		// Echo the message back
		if err := conn.WriteMessage(messageType, message); err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}

// Load Balancer for Reverse Proxy
type LoadBalancer struct {
	servers []string // List of backend servers
	mu      sync.Mutex
	index   int      // Current index for round-robin
}

// Initialize the Load Balancer
func NewLoadBalancer(servers []string) *LoadBalancer {
	return &LoadBalancer{servers: servers}
}

// Get the next server using round-robin
func (lb *LoadBalancer) GetNextServer() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	server := lb.servers[lb.index]
	lb.index = (lb.index + 1) % len(lb.servers)
	return server
}

// Reverse Proxy Handler with Load Balancing
func reverseProxyHandler(lb *LoadBalancer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the next backend server
		targetServer := lb.GetNextServer()
		fmt.Printf("Proxying request to: %s\n", targetServer)
		// Create a new request to the target server
		proxyRequest, err := http.NewRequest(r.Method, targetServer+r.URL.Path, r.Body)
		if err != nil {
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}

		// Copy headers from the original request to the proxy request
		proxyRequest.Header = r.Header

		// Send the request to the target server
		client := &http.Client{}
		proxyResponse, err := client.Do(proxyRequest)
		if err != nil {
			http.Error(w, "Failed to reach the target server", http.StatusBadGateway)
			return
		}
		defer proxyResponse.Body.Close()

		// Copy the headers from the proxy response to the original response
		for key, values := range proxyResponse.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}

		// Write the status code and body back to the original response
		w.WriteHeader(proxyResponse.StatusCode)
		io.Copy(w, proxyResponse.Body)
	}
}

func main() {
	// Define flags for the port number and help
	port := flag.String("port", "8080", "Port to run the server on")
	help := flag.Bool("help", false, "Display help information")
	flag.Parse()

	// Display help information if the help flag is set
	if *help {
		fmt.Println("Usage of the server:")
		flag.PrintDefaults()
		return
	}

	// List of backend servers for load balancing
	backendServers := []string{
		"http://localhost:8080/h1", // Backend server 1
		"http://localhost:8081/h2", // Backend server 2
		// "http://localhost:8080/h3", // Backend server 3
	}

	// Create a load balancer instance
	lb := NewLoadBalancer(backendServers)

	mux := http.NewServeMux()

	// HTTP and WebSocket routes
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/h1", h1)
	mux.HandleFunc("/h2", h2)
	mux.HandleFunc("/h3", h3)

	mux.HandleFunc("/ws", websocketHandler) // Route for WebSocket connection

	// Reverse proxy route with load balancing
	mux.HandleFunc("/proxy", reverseProxyHandler(lb))

	// Wrap the mux with the CORS middleware
	handlerWithCors := corsMiddleware(mux)

	fmt.Printf("Starting server on :%s...\n", *port)
	if err := http.ListenAndServe(":"+*port, handlerWithCors); err != nil {
		fmt.Println("Error starting server:", err)
	}
}