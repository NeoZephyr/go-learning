package handler

import (
	"cron/src/meta"
	"cron/src/util"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("static/hello.html")

		if err != nil {
			io.WriteString(w, "internal server error")
			return
		}

		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		file, head, err := r.FormFile("file")

		if err != nil {
			fmt.Printf("failed to get data, error: %s\n", err.Error())
			return
		}

		defer file.Close()

		fileMeta := meta.FileMeta{
			Name: head.Filename,
			Location: "/tmp/" + head.Filename,
			DateCreated: time.Now().Format("2020-01-01 08:00:00"),
		}

		newFile, err := os.Create(fileMeta.Location)

		if err != nil {
			fmt.Printf("failed to create file, error: %s\n", err.Error())
			return
		}
		defer newFile.Close()
		fileMeta.Size, err = io.Copy(newFile, file)

		if err != nil {
			fmt.Printf("failed to copy file, error: %s\n", err.Error())
			return
		}

		newFile.Seek(0, 0)
		fileMeta.Sha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)
		http.Redirect(w, r, "/file/upload/ok", http.StatusFound)
	}
}

func UploadOkHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload ok")
}