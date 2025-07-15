package httputils

import (
	"net/http"
	"sync"
	"time"
	"golang.org/x/time/rate"
)

// HttpClient represents an HTTP request.
type HttpClient struct {
	client *http.Client
	headers map[string]string
	timeout time.Duration
	mu sync.Mutex
	rateLimiter *rate.Limiter
}

// Response represents an HTTP response.
type Response struct {
	StatusCode  int `json:"status_code"`  // HTTP status code of the response
	Body        []byte `json:"body"`         // Response body as a byte array
	Headers     map[string][]string `json:"headers"`      // Response headers
	Method      string `json:"method"`       // HTTP method used for the request
	URL         string `json:"url"`          // URL of the request
	StartTime   time.Time `json:"start_time"`    // Time the response was received
	ElapsedTime time.Duration `json:"elapsed_time"` // Time taken to get the response
}

// NewHttpClient creates a new instance of HttpClient.
func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: &http.Client{},
		headers: map[string]string{
			"Content-Type": 	"application/json",
			"Accept": 			"application/json",
			"User-Agent":      	"go-nopain/HttpClient", 
			"Accept-Encoding": 	"gzip, deflate",
			"Cache-Control":   	"no-cache",
			"Connection":      	"keep-alive",
		},
		timeout: 10 * time.Second,
		rateLimiter: nil,
	}
}
