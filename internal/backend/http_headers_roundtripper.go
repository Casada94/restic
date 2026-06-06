package backend

import "net/http"

// httpHeadersRoundTripper injects a set of headers into outgoing HTTP requests.
type httpHeadersRoundTripper struct {
	headers map[string]string
	rt      http.RoundTripper
}

func newCustomHeadersRoundTripper(rt http.RoundTripper, headers map[string]string) *httpHeadersRoundTripper {
	return &httpHeadersRoundTripper{rt: rt, headers: headers}
}

func (c *httpHeadersRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}
	return c.rt.RoundTrip(req)
}
