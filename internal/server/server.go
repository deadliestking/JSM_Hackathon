package server

import (
	"fmt"
	"net/http"
	"io"
	"os"
	"path/filepath"
	"time"
)

// Server represents a CDN edge server
type Server struct {
	ID      int
	Address string
	Cache   map[string]string
	Load    int
	RootDir string  // Directory to store files
}

// ServeHTTP simulates serving content from a CDN server
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Load++
	url := r.URL.Path

	// Handle file uploads
	if r.Method == http.MethodPost && url == "/upload" {
		s.UploadHandler(w, r)
		s.Load--
		return
	}

	// Serve files under the "/files" route
	if url == "/files" || url[:7] == "/files/" {
		filePath := filepath.Join(s.RootDir, url[7:])  // Get the actual file path
	
		// Check if file exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
	
		// Set headers to force file download
		w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Transfer-Encoding", "binary")
		w.Header().Set("Expires", "0")
	
		// Serve the file
		http.ServeFile(w, r, filePath)
		s.Load--
		return
	}
	
	

	// Handle file serving and caching
	if r.Header.Get("Cache-Control") == "no-cache" {
		delete(s.Cache, url)
		content := fetchFromOrigin(url)
		s.Cache[url] = content
		fmt.Fprintf(w, "Cache busted, fetched fresh content on server %d: %s\n", s.ID, content)
	} else if content, exists := s.Cache[url]; exists {
		fmt.Fprintf(w, "Serving from cache on server %d: %s\n", s.ID, content)
	} else {
		content := fetchFromOrigin(url)
		s.Cache[url] = content
		fmt.Fprintf(w, "Fetched from origin and cached on server %d: %s\n", s.ID, content)
	}

	s.Load--
}

// UploadHandler handles file uploads
func (s *Server) UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10MB max upload size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusInternalServerError)
		return
	}

	// Get the file from form data
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Save the file to the root directory
	filePath := filepath.Join(s.RootDir, handler.Filename)
	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}
	defer out.Close()
	io.Copy(out, file)

	// Update cache with the new file path
	s.Cache["/"+handler.Filename] = filePath

	fmt.Fprintf(w, "File uploaded successfully: %s\n", handler.Filename)
}

// Simulate fetching from the origin server
func fetchFromOrigin(url string) string {
	time.Sleep(100 * time.Millisecond) // Simulate network latency
	return "Content of " + url
}
