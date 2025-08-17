package pastebin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Package pastebin contains a generic HTTP client for a "paste" provider.
// We intentionally keep it provider-agnostic and assume a common JSON shape.
//
// Assumed API (adjust to your provider):
//   POST {baseURL}/pastes
//   Body: { "title": "...", "content": "...", "private": true, "expires": "10m|1h|1d|never" }
//   201 -> { "id": "abc123", "url": "https://.../abc123" }
//
// If your provider differs, change `CreatePaste` in one place.

// Client is safe for Concurrent user if underlying http.client is safe (it is).
type Client struct {
	baseURL string
	apiKey  string
	http    *http.Client
}

// New Constructs a new client. Pass a custom http.Client to control timeouts,
// proxies, and transport. If nil, a default client with timeout is used.
func New(baseURL, apiKey string, httpClient *http.Client, defaultTimeout time.Duration) *Client {

	// create http client if nil
	if httpClient == nil {
		httpClient = &http.Client{Timeout: defaultTimeout}
	}

	// return client
	return &Client{
		baseURL: strings.TrimRight(baseURL, "/"),
		apiKey:  apiKey,
		http:    httpClient,
	}
}

// CreatePasteRequest describes what we send upstream
type CreatePasteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Private bool   `json:"private"`
	// "expires" is provider-defined. We keep a string to be flexible (e.g. "10m", "1h", "never").
	Expires string `json:"expires"`
}

// CreatePasteResponse is a minimal success shape.
type CreatePasteResponse struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

// CreatePaste posts a new paste with context cancellation.
// We always set Content-Type and Authorization (if key present).
func (c *Client) CreatePaste(ctx context.Context, req CreatePasteRequest) (*CreatePasteResponse, error) {
	if c.baseURL == "" {
		return nil, fmt.Errorf("pastebin: base URL is empty; set API_BASE_URL or pass --api-base-url")
	}

	b, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/pastes", bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if c.apiKey != "" {
		httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		// We include body snippet to aid debugging; truncate in real-world if body can be huge.
		return nil, fmt.Errorf("upstream error: %s; body: %s", resp.Status, string(body))
	}

	var out CreatePasteResponse
	if err := json.Unmarshal(body, &out); err != nil {
		// If provider returns plain text (some do), surface it.
		return nil, fmt.Errorf("decode response: %w, raw: %s", err, string(body))
	}

	if out.URL == "" && out.ID == "" {
		return nil, fmt.Errorf("unexpected response shape; got; %s", string(body))
	}

	return &out, nil
}
