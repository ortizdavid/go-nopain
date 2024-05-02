package httputils

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/ortizdavid/go-nopain/serialization"
)

// HttpClient represents an HTTP request.
type HttpClient struct {
	client *http.Client
	headers map[string]string
}

// Response represents an HTTP response.
type Response struct {
	StatusCode int                `json:"statusCode"`             // HTTP status code of the response
	Body       []byte             `json:"body"`                   // Response body as a byte array
	Headers    map[string][]string `json:"headers"`                // Response headers
}


// NewHttpClient creates a new instance of HttpClient.
func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: &http.Client{},
		headers: make(map[string]string),
	}
}


// SetHeader sets a custom header on the HTTP client.
func (cl *HttpClient) SetHeader(key string, value string) {
	cl.headers[key] = value
}

//GetHeader return all headers on the GTTP client
func (cl *HttpClient) GetHeader(key string) string {
	return cl.headers[key]
}

//GetAllHeaders return all headers on the HTTP client
func (cl *HttpClient) GetAllHeaders() map[string]string {
	return cl.headers
}

// Get performs an HTTP GET request to the specified URL with custom headers and returns a Response.
func (cl *HttpClient) Get(url string) (*Response, error) {
	// Create a new HTTP GET request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// Set custom headers on the request
	if cl.headers != nil && len(cl.headers) > 0 {
		for key, value := range cl.headers {
			req.Header.Set(key, value)
		}
	}
	// Perform the HTTP GET request
	resp, err := cl.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Create and return the response object
	return &Response{
		StatusCode: resp.StatusCode,
		Body:       body,
		Headers:    resp.Header,
	}, nil
}


// Post performs an HTTP POST request to the specified URL with the provided data and returns a Response.
func (cl *HttpClient) Post(url string, data interface{}) (*Response, error) {
	// get Body reader
	bodyReader, err := cl.getBodyReader(data)
	if err != nil {
		return nil, err
	}
	// Create a new HTTP POST request
	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return nil, err
	}
	// Set custom headers on the request if provided
	if cl.headers != nil && len(cl.headers) > 0 {
		for key, value := range cl.headers {
			req.Header.Set(key, value)
		}
	}
	// Perform the HTTP POST request
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
	}, nil
}


// Put performs an HTTP PUT request to the specified URL with the provided data and returns a Response.
func (cl *HttpClient) Put(url string, data interface{}) (*Response, error) {
	// get Body reader
	bodyReader, err := cl.getBodyReader(data)
	if err != nil {
		return nil, err
	}
	// Create a new HTTP PUT request
	req, err := http.NewRequest(http.MethodPut, url, bodyReader)
	if err != nil {
		return nil, err
	}
	// Set custom headers on the request if provided
	if cl.headers != nil && len(cl.headers) > 0 {
		for key, value := range cl.headers {
			req.Header.Set(key, value)
		}
	}
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
	}, nil
}


// Delete performs an HTTP DELETE request to the specified URL and returns a Response.
func (cl *HttpClient) Delete(url string) (*Response, error) {
	// Create a new HTTP Delete request
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	// Set custom headers on the request
	if cl.headers != nil && len(cl.headers) > 0 {
		for key, value := range cl.headers {
			req.Header.Set(key, value)
		}
	}
	// Perform the HTTP Delete request
	resp, err := cl.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Create and return the response object
	return &Response{
		StatusCode: resp.StatusCode,
		Body:       body,
		Headers:    resp.Header,
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