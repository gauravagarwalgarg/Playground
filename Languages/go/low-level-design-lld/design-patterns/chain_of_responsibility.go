package main

import "fmt"

// Chain of Responsibility Pattern: Passes a request along a chain of handlers.
// Each handler decides to process the request or pass it to the next handler.
// Example: HTTP middleware / request validation pipeline.

type Request struct {
	UserID    string
	Token     string
	Body      string
	IP        string
	RateCount int
}

type Response struct {
	Code    int
	Message string
}

// Handler interface
type Handler interface {
	SetNext(h Handler) Handler
	Handle(req *Request) *Response
}

// BaseHandler provides default chaining behavior
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(h Handler) Handler {
	b.next = h
	return h
}

func (b *BaseHandler) HandleNext(req *Request) *Response {
	if b.next != nil {
		return b.next.Handle(req)
	}
	return &Response{Code: 200, Message: "OK"}
}

// AuthHandler - checks authentication
type AuthHandler struct {
	BaseHandler
	validTokens map[string]bool
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		validTokens: map[string]bool{"token-abc": true, "token-xyz": true},
	}
}

func (h *AuthHandler) Handle(req *Request) *Response {
	if req.Token == "" {
		return &Response{Code: 401, Message: "missing auth token"}
	}
	if !h.validTokens[req.Token] {
		return &Response{Code: 403, Message: "invalid token"}
	}
	fmt.Println("[Auth] ✓ token valid")
	return h.HandleNext(req)
}

// RateLimitHandler - checks rate limiting
type RateLimitHandler struct {
	BaseHandler
	maxRequests int
}

func NewRateLimitHandler(max int) *RateLimitHandler {
	return &RateLimitHandler{maxRequests: max}
}

func (h *RateLimitHandler) Handle(req *Request) *Response {
	if req.RateCount > h.maxRequests {
		return &Response{Code: 429, Message: "rate limit exceeded"}
	}
	fmt.Println("[RateLimit] ✓ within limits")
	return h.HandleNext(req)
}

// ValidationHandler - validates request body
type ValidationHandler struct {
	BaseHandler
}

func (h *ValidationHandler) Handle(req *Request) *Response {
	if req.Body == "" {
		return &Response{Code: 400, Message: "request body is empty"}
	}
	fmt.Println("[Validation] ✓ body valid")
	return h.HandleNext(req)
}

// LoggingHandler - logs the request
type LoggingHandler struct {
	BaseHandler
}

func (h *LoggingHandler) Handle(req *Request) *Response {
	fmt.Printf("[Logging] user=%s ip=%s\n", req.UserID, req.IP)
	return h.HandleNext(req)
}

func main() {
	// Build the chain: Logging → Auth → RateLimit → Validation
	logging := &LoggingHandler{}
	auth := NewAuthHandler()
	rateLimit := NewRateLimitHandler(100)
	validation := &ValidationHandler{}

	logging.SetNext(auth).SetNext(rateLimit).SetNext(validation)

	// Test 1: Valid request passes all handlers
	fmt.Println("--- Test 1: Valid request ---")
	resp := logging.Handle(&Request{
		UserID: "user1", Token: "token-abc",
		Body: `{"action":"create"}`, IP: "192.168.1.1", RateCount: 5,
	})
	if resp.Code == 200 {
		fmt.Println("PASS: valid request, code:", resp.Code)
	} else {
		panic(fmt.Sprintf("FAIL: expected 200, got %d", resp.Code))
	}

	// Test 2: Missing token
	fmt.Println("\n--- Test 2: Missing token ---")
	resp = logging.Handle(&Request{UserID: "user2", Body: "data", IP: "10.0.0.1"})
	if resp.Code == 401 {
		fmt.Println("PASS: rejected, code:", resp.Code, resp.Message)
	} else {
		panic("FAIL: expected 401")
	}

	// Test 3: Rate limit exceeded
	fmt.Println("\n--- Test 3: Rate limit ---")
	resp = logging.Handle(&Request{
		UserID: "user3", Token: "token-xyz",
		Body: "data", IP: "10.0.0.2", RateCount: 150,
	})
	if resp.Code == 429 {
		fmt.Println("PASS: rate limited, code:", resp.Code)
	} else {
		panic("FAIL: expected 429")
	}

	// Test 4: Empty body
	fmt.Println("\n--- Test 4: Empty body ---")
	resp = logging.Handle(&Request{
		UserID: "user4", Token: "token-abc", IP: "10.0.0.3", RateCount: 1,
	})
	if resp.Code == 400 {
		fmt.Println("PASS: validation failed, code:", resp.Code)
	} else {
		panic("FAIL: expected 400")
	}
}
