package step5

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"
)

type APIResponse struct {
	URL        string
	Data       string
	StatusCode int
	Err        error
}

func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	results := make([]*APIResponse, len(urls))
	var wg sync.WaitGroup

	for i, url := range urls {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()

			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				results[i] = &APIResponse{
					URL: url,
					Err: err,
				}
				return
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				results[i] = &APIResponse{
					URL: url,
					Err: err,
				}
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				results[i] = &APIResponse{
					URL: url,
					Err: err,
				}
				return
			}

			results[i] = &APIResponse{
				URL:        url,
				Data:       string(body),
				StatusCode: resp.StatusCode,
				Err:        nil,
			}
		}(i, url)
	}

	wg.Wait()

	return results
}
