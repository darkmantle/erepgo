package erepgo

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const baseURL = "https://api.erepublik.com"

// Client is an authenticated eRepublik API client.
// See https://api.erepublik.com/doc/
type Client struct {
	PublicKey  string
	secretKey  string
	httpClient *http.Client
	format     string // "application/json" or "application/xml"
}

// NewClient returns a new eRepublik API client.
// publicKey and secretKey are the API keys from eRepublik.
func NewClient(publicKey, secretKey string) *Client {
	return &Client{
		PublicKey:  publicKey,
		secretKey:  secretKey,
		httpClient: &http.Client{Timeout: 30 * time.Second},
		format:     "application/json",
	}
}

// SetFormat sets the response format: "json" (default) or "xml".
func (c *Client) SetFormat(format string) {
	switch strings.ToLower(format) {
	case "xml":
		c.format = "application/xml"
	default:
		c.format = "application/json"
	}
}

// digest builds the HMAC-SHA256 digest for the request.
// Message is: lower(resource:action[:params]):Date (RFC 1123).
func (c *Client) digest(resource, action string, params url.Values) (string, string) {
	date := time.Now().UTC().Format(time.RFC1123)
	resource = strings.ToLower(resource)
	action = strings.ToLower(action)

	var msg string
	if len(params) == 0 {
		msg = strings.ToLower(resource+":"+action) + ":" + date
	} else {
		encoded := params.Encode()
		msg = strings.ToLower(resource+":"+action+":"+encoded) + ":" + date
	}

	mac := hmac.New(sha256.New, []byte(c.secretKey))
	mac.Write([]byte(msg))
	digest := hex.EncodeToString(mac.Sum(nil))
	return digest, date
}

// Call performs an authenticated request to resource/action with optional query params.
// Returns the response body and error. Use RawCall for raw bytes, or decode JSON manually.
func (c *Client) Call(resource, action string, params map[string]string) ([]byte, error) {
	return c.RawCall(resource, action, params)
}

// RawCall performs an authenticated request and returns the response body.
func (c *Client) RawCall(resource, action string, params map[string]string) ([]byte, error) {
	u, err := url.Parse(fmt.Sprintf("%s/%s/%s", baseURL, resource, action))
	if err != nil {
		return nil, err
	}

	q := make(url.Values)
	for k, v := range params {
		q.Set(k, v)
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	digest, dateStr := c.digest(resource, action, q)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Date", dateStr)
	req.Header.Set("Auth", c.PublicKey+"/"+digest)
	req.Header.Set("Accept", c.format)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api error: %s (%d): %s", resp.Status, resp.StatusCode, string(body))
	}
	return body, nil
}

// CallJSON performs the request and decodes the JSON response into v.
func (c *Client) CallJSON(resource, action string, params map[string]string, v interface{}) error {
	body, err := c.RawCall(resource, action, params)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}
