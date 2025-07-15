package httputils

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

// doRequest performs an HTTP request with the provided method, URL, and data.
func (cl *HttpClient) doRequest(ctx context.Context, url string, method string, data interface{}) (*Response, error) {
	startTime := time.Now()
	// Wait for rate limiter
	if cl.rateLimiter != nil {
		if err := cl.rateLimiter.Wait(ctx); err != nil {
			return nil, err
		}
	}
	//Get Body Reader
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
	// Calculate the elapsed time
	elapsedTime := time.Since(startTime)
	// Create and return the response object
	return &Response{
		StatusCode:  resp.StatusCode,
		Body:        respBody,
		Headers:     resp.Header,
		Method:      method,
		URL:         url,
		StartTime:   startTime,
		ElapsedTime: elapsedTime,
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
	case io.Reader:
		bodyReader = d
	default:
		// If data is of any other type, try to encode it as JSON and convert to a reader
		jsonData, err := json.Marshal(d)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(jsonData)
	}
	return bodyReader, nil
}


