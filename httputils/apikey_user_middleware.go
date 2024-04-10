package httputils

import "sync"

type ApiKeyUserMiddleware struct {
	mu sync.Mutex
}

// Return a ApiKeyUserMiddleware object with the apiKey passed
func NewApiKeyUserMiddleWare(apiKey string) ApiKeyMiddleware {
	return ApiKeyMiddleware{
		defaultKey: apiKey,
	}
}