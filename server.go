package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler for the root path
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go Server!\n")
	fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
}

// Handler for /hello path
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

// Handler for /health path
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK\n")
}

// Handler for logging and navigation page
func logNavHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
	<title>Go Server - Logging and Navigation</title>
</head>
<body>
	<h1>Go Server - Logging and Navigation</h1>
	<ul>
		<li><a href="/">Home</a></li>
		<li><a href="/hello">Hello</a></li>
		<li><a href="/health">Health Check</a></li>
		<li><a href="/logs">View Logs</a></li>
	</ul>
	<div id="logs">
		<h2>Server Logs</h2>
		<pre id="log-output">Log output will appear here...</pre>
	</div>
</body>
</html>
`)
}

// Handler for logs endpoint
func logsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "No logs available yet.\n")
}

func main() {
	// Register handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/lognav", logNavHandler)
	http.HandleFunc("/logs", logsHandler)

	// Define port
	port := ":8080"

	// Start server
	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}