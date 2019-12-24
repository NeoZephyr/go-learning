package mock

type Downloader struct {
	Content string
}

func (downloader *Downloader) Download(url string) string {
	return downloader.Content
}