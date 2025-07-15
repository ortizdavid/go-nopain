package httputils

import (
	"context"
	"net/http"
)

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
