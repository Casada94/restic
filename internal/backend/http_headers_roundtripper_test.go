package backend

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCustomHeadersTransport(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("X-Backup-Token"); got != "abc123" {
			t.Fatalf("expected X-Backup-Token header abc123, got %q", got)
		}
		if got := r.Header.Get("X-Request-Source"); got != "restic" {
			t.Fatalf("expected X-Request-Source header restic, got %q", got)
		}
		w.WriteHeader(http.StatusOK)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	transport := &httpHeadersRoundTripper{
		headers: map[string]string{
			"X-Backup-Token":   "abc123",
			"X-Request-Source": "restic",
		},
		rt: http.DefaultTransport,
	}

	client := &http.Client{Transport: transport}
	resp, err := client.Get(server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			t.Log("failed to close response body")
		}
	}()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
