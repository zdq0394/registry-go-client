package registry

import (
	"crypto/tls"
	"net/http"
	"time"
)

func getDefaultHTTPClient() *http.Client {
	t := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
	}
	client := &http.Client{
		Transport: t,
		Timeout:   30 * time.Second,
	}
	return client
}

func getBearerTokenClient(token string) *http.Client {
	return &http.Client{
		Transport: NewBearerTokenTransport(token, http.DefaultTransport),
	}
}

type BearerTokenTransport struct {
	bearerToken string
	transport   http.RoundTripper
}

func NewBearerTokenTransport(token string, trans http.RoundTripper) http.RoundTripper {
	if trans == nil {
		trans = http.DefaultTransport
	}
	return &BearerTokenTransport{
		bearerToken: token,
		transport:   trans,
	}
}

func (p *BearerTokenTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	req.Header.Set("Authorization", "Bearer "+p.bearerToken)
	resp, err = p.transport.RoundTrip(req)
	return
}
