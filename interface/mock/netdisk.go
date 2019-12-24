package mock

import "fmt"

type NetDisk struct {
	Content string
}

func (netDisk *NetDisk) Download(url string) string {
	return netDisk.Content
}

func (netDisk *NetDisk) Upload(url string, form map[string]string) string {
	netDisk.Content = form["content"]
	return "ok"
}

// Stringer 接口
func (netDisk *NetDisk) String() string {
	return fmt.Sprintf("mock.NetDisk, {Content: %s}", netDisk.Content)
}