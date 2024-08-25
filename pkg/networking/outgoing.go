package networking

import (
	"fmt"
	"io"
	"net/http"
)

func MakeRequest(client *http.Client, method string, url string, body io.Reader, headers *map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("error formatting request: %s", err)
	}

	for k, v := range *headers {
		req.Header.Add(k, v)
	}

	return client.Do(req)
}
