package httputils

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
	"github.com/ortizdavid/go-nopain/serialization"
)

// HttpClient represents an HTTP request.
type HttpClient struct {
	client *http.Client
	headers map[string]string
	timeout time.Duration
	mu sync.Mutex
}

// Response represents an HTTP response.
type Response struct {
	StatusCode int `json:"status_code"`// HTTP status code of the response
	Body       []byte `json:"body"` // Response body as a byte array
	Headers    map[string][]string`json:"headers"` // Response headers
	TimesTamp time.Time
}

// NewHttpClient creates a new instance of HttpClient.
func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: &http.Client{},
		headers: map[string]string{
			"Content-Type": "application/json",
		},
		timeout: 10 * time.Second,
	}
}

// SetHeader sets a custom header on the HTTP client.
func (cl *HttpClient) SetHeader(key string, value string) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.headers[key] = value
}

//GetHeader return all headers on the GTTP client
func (cl *HttpClient) GetHeader(key string) (string, bool) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	val, ok := cl.headers[key]
	return val, ok
}

//GetAllHeaders return all headers on the HTTP client
func (cl *HttpClient) GetAllHeaders() map[string]string {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	headerCopy := make(map[string]string, len(cl.headers))
	for k, v := range cl.headers {
		headerCopy[k] = v
	}
	return headerCopy
}

// RemoveHeader removes a custom header from the HTTP client.
func (cl *HttpClient) RemoveHeader(key string) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	delete(cl.headers, key)
}

// SetHeader sets timout for request
func (cl *HttpClient) SetTimeout(newTimeout time.Duration) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.timeout = newTimeout
}

// Get performs an HTTP GET request to the specified URL with custom headers and returns a Response.
func (cl *HttpClient) Get(url string) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cl.timeout)
	defer cancel()
	return cl.doRequest(ctx, url, http.MethodGet, nil)
}

// Post performs an HTTP POST request to the specified URL with the provided data and returns a Response.
func (cl *HttpClient) Post(url string, data interface{}) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cl.timeout)
	defer cancel()
	return cl.doRequest(ctx, url, http.MethodPost, data)
}

// Put performs an HTTP PUT request to the specified URL with the provided data and returns a Response.
func (cl *HttpClient) Put(url string, data interface{}) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cl.timeout)
	defer cancel()
	return cl.doRequest(ctx, url, http.MethodPut, data)
}

// Delete performs an HTTP DELETE request to the specified URL and returns a Response.
func (cl *HttpClient) Delete(url string) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cl.timeout)
	defer cancel()
	return cl.doRequest(ctx, url, http.MethodDelete, nil)
}

// Patch performs an HTTP PATCH request to the specified URL with the provided data.
func (cl *HttpClient) Patch(url string, data interface{}) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cl.timeout)
	defer cancel()
	return cl.doRequest(ctx, http.MethodPatch, url, data)
}

// Options performs an HTTP OPTIONS request to the specified URL.
func (cl *HttpClient) Options(url string) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cl.timeout)
	defer cancel()
	return cl.doRequest(ctx, http.MethodOptions, url, nil)
}

// Head performs an HTTP HEAD request to the specified URL.
func (cl *HttpClient) Head(url string) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cl.timeout)
	defer cancel()
	return cl.doRequest(ctx, http.MethodHead, url, nil)
}

// doRequest performs an HTTP request with the provided method, URL, and data.
func (cl *HttpClient) doRequest(ctx context.Context, url string, method string, data interface{}) (*Response, error) {
	bodyReader, err := cl.getBodyReader(data)
	if err != nil {
		return nil, err
	}
	// Create a new HTTP PUT request
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, err
	}
	cl.mu.Lock()
	// Set custom headers on the request if provided
	if cl.headers != nil && len(cl.headers) > 0 {
		for key, value := range cl.headers {
			req.Header.Set(key, value)
		}
	}
	cl.mu.Unlock()
	// Perform the HTTP PUT request
	resp, err := cl.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Create and return the response object
	return &Response{
		StatusCode: resp.StatusCode,
		Body:       respBody,
		Headers:    resp.Header,
		TimesTamp: time.Now(),
	}, nil
}

// getBodyReader returns an io.Reader for the provided data.
// It converts the data to a reader based on its type.
func (cl *HttpClient) getBodyReader(data interface{}) (io.Reader, error) {
	var bodyReader io.Reader

	switch d := data.(type) {
	case string:
		// If data is already a string, convert it to a reader
		bodyReader = strings.NewReader(d)
	case []byte:
		// If data is a byte slice, convert it to a reader
		bodyReader = bytes.NewReader(d)
	default:
		// If data is of any other type, try to encode it as JSON and convert to a reader
		jsonData, err := serialization.SerializeJson(d)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(jsonData)
	}
	return bodyReader, nil
}


