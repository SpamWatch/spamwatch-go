package spamwatch

import (
	"fmt"
	"io"
	"net/http"
)

func Client(token string, opts *ClientOpts) *SpamWatch {
	if opts == nil {
		opts = &ClientOpts{}
	}
	if string(opts.ApiUrl[len(opts.ApiUrl)-1]) != "/" {
		opts.ApiUrl = opts.ApiUrl + "/"
	}
	if opts.ApiUrl == "" {
		opts.ApiUrl = DEFAULT_API_URL
	}
	var sw = &SpamWatch{}
	sw.TimeOut = opts.TimeOut
	sw.Token.Token = token
	sw.ApiUrl = opts.ApiUrl
	// Ignore error on GetSelf method as it would already be handled in succeeding calls
	t, _ := sw.GetSelf()
	if t != nil {
		sw.Token = *t
	}
	return sw
}

func (s *SpamWatch) MakeRequest(http_method string, method string, body io.Reader) ([]byte, error) {
	r, err := http.NewRequest(http_method, s.ApiUrl+method, body)
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Token.Token))
	if body != nil {
		r.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to build %s request to %s: %w", http_method, method, err)
	}
	var client = http.Client{
		Timeout: s.TimeOut,
	}

	resp, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("failed to execute %s request to %s: %w", http_method, method, err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = handleError(s, resp.StatusCode, b, method)
	if err != nil {
		return nil, err
	}
	return b, nil
}
