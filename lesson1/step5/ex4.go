package step5

import (
	"io"
	"net/http"
	"time"
)

func StartServer(maxTimeout time.Duration) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:8081/provideData")
		if err != nil {
			http.Error(w, "", http.StatusServiceUnavailable)
			return
		}
		defer resp.Body.Close()

		io.Copy(w, resp.Body)
	})

	timeoutHandler := http.TimeoutHandler(handler, maxTimeout, "")

	http.Handle("/readSource", timeoutHandler)

	http.ListenAndServe(":8080", nil)
}
