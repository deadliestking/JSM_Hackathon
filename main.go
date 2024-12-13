package main

import (
	"cdn-simulator/internal/loadbalancer"
	"cdn-simulator/internal/server"
	"fmt"
	"net/http"
)

func main() {
	// Create CDN servers
	server1 := &server.Server{ID: 1, Address: "localhost:8081", Cache: make(map[string]string), RootDir: "./server1_files"}
	server2 := &server.Server{ID: 2, Address: "localhost:8082", Cache: make(map[string]string), RootDir: "./server2_files"}

	// Create a load balancer
	lb := &loadbalancer.LoadBalancer{
		Servers: []*server.Server{server1, server2},
	}

	// Start the CDN servers
	go func() {
		fmt.Println("Server 1 is running on port 8081...")
		http.ListenAndServe(":8081", server1)
	}()
	go func() {
		fmt.Println("Server 2 is running on port 8082...")
		http.ListenAndServe(":8082", server2)
	}()

	// Start the load balancer
	fmt.Println("CDN Load Balancer is running on port 8080...")
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Pass the file path (URL) to the load balancer
		server, err := lb.GetServer(r.URL.Path)
		if err != nil {
			http.Error(w, "No servers available", http.StatusServiceUnavailable)
			return
		}
		server.ServeHTTP(w, r)
	}))
}
