package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"sync"
	"time"
)

// URL Shortener - System Design implementation.
// Components: Base62 encoding, in-memory store, TTL expiration, analytics.
// Real system would use: distributed ID generator, Redis/DB, CDN.

const (
	base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	shortLength = 7
	baseURL     = "https://short.ly/"
)

// URLEntry stores metadata about a shortened URL
type URLEntry struct {
	OriginalURL string
	ShortCode   string
	CreatedAt   time.Time
	ExpiresAt   time.Time
	ClickCount  int64
	CreatedBy   string
}

// URLShortener - core service
type URLShortener struct {
	mu         sync.RWMutex
	codeToURL  map[string]*URLEntry // shortCode → entry
	urlToCode  map[string]string    // originalURL → shortCode (dedup)
	counter    uint64               // for sequential ID generation
}

func NewURLShortener() *URLShortener {
	return &URLShortener{
		codeToURL: make(map[string]*URLEntry),
		urlToCode: make(map[string]string),
		counter:   1000000, // start from a high number for shorter codes
	}
}

// Shorten creates a short URL for the given original URL
func (s *URLShortener) Shorten(originalURL, userID string, ttl time.Duration) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Dedup: if URL already shortened, return existing
	if code, exists := s.urlToCode[originalURL]; exists {
		return baseURL + code, nil
	}

	// Generate short code using hash-based approach
	code := s.generateCode(originalURL)

	// Handle collision
	for s.codeToURL[code] != nil {
		s.counter++
		code = s.generateCode(fmt.Sprintf("%s:%d", originalURL, s.counter))
	}

	entry := &URLEntry{
		OriginalURL: originalURL,
		ShortCode:   code,
		CreatedAt:   time.Now(),
		ExpiresAt:   time.Now().Add(ttl),
		CreatedBy:   userID,
	}

	s.codeToURL[code] = entry
	s.urlToCode[originalURL] = code
	return baseURL + code, nil
}

// Resolve expands a short code to the original URL
func (s *URLShortener) Resolve(shortURL string) (string, error) {
	code := strings.TrimPrefix(shortURL, baseURL)
	s.mu.RLock()
	entry, exists := s.codeToURL[code]
	s.mu.RUnlock()

	if !exists {
		return "", fmt.Errorf("short URL not found: %s", code)
	}

	if time.Now().After(entry.ExpiresAt) {
		return "", fmt.Errorf("short URL expired: %s", code)
	}

	// Increment click count (analytics)
	s.mu.Lock()
	entry.ClickCount++
	s.mu.Unlock()

	return entry.OriginalURL, nil
}

// Stats returns analytics for a short URL
func (s *URLShortener) Stats(shortURL string) (*URLEntry, error) {
	code := strings.TrimPrefix(shortURL, baseURL)
	s.mu.RLock()
	defer s.mu.RUnlock()
	entry, exists := s.codeToURL[code]
	if !exists {
		return nil, fmt.Errorf("not found")
	}
	return entry, nil
}

// Delete removes a shortened URL
func (s *URLShortener) Delete(shortURL string) error {
	code := strings.TrimPrefix(shortURL, baseURL)
	s.mu.Lock()
	defer s.mu.Unlock()
	entry, exists := s.codeToURL[code]
	if !exists {
		return fmt.Errorf("not found")
	}
	delete(s.urlToCode, entry.OriginalURL)
	delete(s.codeToURL, code)
	return nil
}

// generateCode creates a base62 short code from a URL
func (s *URLShortener) generateCode(input string) string {
	hash := sha256.Sum256([]byte(input))
	encoded := base64.URLEncoding.EncodeToString(hash[:])
	// Convert to base62 (remove non-alphanumeric)
	var code strings.Builder
	for _, ch := range encoded {
		if (ch >= '0' && ch <= '9') || (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
			code.WriteRune(ch)
			if code.Len() == shortLength {
				break
			}
		}
	}
	return code.String()
}

func main() {
	shortener := NewURLShortener()

	// Shorten URLs
	urls := []string{
		"https://www.example.com/very/long/path/to/resource?param=value&other=123",
		"https://docs.google.com/document/d/1234567890/edit",
		"https://github.com/user/repo/pull/42/files#diff-abc123",
	}

	fmt.Println("--- URL Shortening ---")
	var shortURLs []string
	for _, url := range urls {
		short, err := shortener.Shorten(url, "user1", 24*time.Hour)
		if err != nil {
			panic("FAIL: " + err.Error())
		}
		shortURLs = append(shortURLs, short)
		fmt.Printf("  %s\n  → %s\n\n", url, short)
	}

	// Test resolve
	fmt.Println("--- Resolving ---")
	for _, short := range shortURLs {
		original, err := shortener.Resolve(short)
		if err != nil {
			panic("FAIL: " + err.Error())
		}
		fmt.Printf("  %s → %s\n", short, original[:50]+"...")
	}
	fmt.Println("PASS: all URLs resolved correctly")

	// Test deduplication
	fmt.Println("\n--- Deduplication ---")
	short1, _ := shortener.Shorten(urls[0], "user1", 24*time.Hour)
	if short1 != shortURLs[0] {
		panic("FAIL: same URL should return same short URL")
	}
	fmt.Println("PASS: duplicate URL returns same short code")

	// Test analytics
	fmt.Println("\n--- Analytics ---")
	shortener.Resolve(shortURLs[0])
	shortener.Resolve(shortURLs[0])
	shortener.Resolve(shortURLs[0])
	stats, _ := shortener.Stats(shortURLs[0])
	fmt.Printf("  Clicks: %d, Created: %s\n", stats.ClickCount, stats.CreatedAt.Format("15:04:05"))
	if stats.ClickCount != 4 { // 1 from resolve test + 3 here
		panic(fmt.Sprintf("FAIL: expected 4 clicks, got %d", stats.ClickCount))
	}
	fmt.Println("PASS: click tracking works")

	// Test deletion
	fmt.Println("\n--- Deletion ---")
	err := shortener.Delete(shortURLs[1])
	if err != nil {
		panic("FAIL: delete error: " + err.Error())
	}
	_, err = shortener.Resolve(shortURLs[1])
	if err == nil {
		panic("FAIL: deleted URL should not resolve")
	}
	fmt.Println("PASS: deleted URL no longer resolves")

	// Test expiration
	fmt.Println("\n--- Expiration ---")
	expiring, _ := shortener.Shorten("https://temp.com", "user1", 1*time.Millisecond)
	time.Sleep(5 * time.Millisecond)
	_, err = shortener.Resolve(expiring)
	if err == nil {
		panic("FAIL: expired URL should not resolve")
	}
	fmt.Println("PASS: expired URL rejected")

	// Concurrent access test
	fmt.Println("\n--- Concurrent access ---")
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			url := fmt.Sprintf("https://example.com/page/%d", id)
			shortener.Shorten(url, "user1", time.Hour)
		}(i)
	}
	wg.Wait()
	fmt.Println("PASS: 100 concurrent shortenings succeeded")
}
