package text

import (
	"net/http"
	"net/http/httputil"
)

type Downloader struct {
	UserAgent string
}

func (downloader *Downloader) Download(url string) string {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(response, true)

	response.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(result)
}
