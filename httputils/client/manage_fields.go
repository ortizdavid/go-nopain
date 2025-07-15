package httputils

import (
	"time"
	"golang.org/x/time/rate"
)

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

// SetRateLimiter sets or updates the rate limiter for the HTTP client
func (cl *HttpClient) SetRateLimiter(rateLimitDuration time.Duration, brust int) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	rateLimit := rate.Every(rateLimitDuration)
	newRateLimiter := rate.NewLimiter(rateLimit, brust)
	cl.rateLimiter = newRateLimiter
}
