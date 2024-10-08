package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/hoomanist/allegro-server/pkg/auth"
	"github.com/hoomanist/allegro-server/pkg/database"
)

const MAX_UPLOAD_SIZE = 120 * 1024 * 1024

type Progress struct {
	TotalSize int64
	BytesRead int64
	FileName  string
	User      string
}

func (pr *Progress) Print() {
	if pr.BytesRead == pr.TotalSize {
		//ToDo: write this to a logging database
		fmt.Printf("%s is Uploaded by %s!", pr.FileName, pr.User)
		return
	}
	// fmt.Printf("File Upload in progress: %d\n", pr.BytesRead)
}

func (pr *Progress) Write(p []byte) (n int, err error) {
	n, err = len(p), nil
	pr.BytesRead += int64(n)
	pr.Print()
	return
}

func (s *server) FileUpload(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	files := r.MultipartForm.File["file"]

	for _, fileHeader := range files {
		if fileHeader.Size > MAX_UPLOAD_SIZE {
			http.Error(w, fmt.Sprintf("The uploaded file is too big: %s. use a file smaller than 20 Mb", fileHeader.Filename), http.StatusBadGateway)
			return
		}
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()
		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = os.MkdirAll("./uploads", os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t := time.Now()
		fpath := fmt.Sprintf("./uploads/%d%s", t.UnixNano(), filepath.Ext(fileHeader.Filename))
		f, err := os.Create(fpath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer f.Close()
		username, err := auth.ExtractClaims(w, r, s.Key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		pr := &Progress{
			TotalSize: fileHeader.Size,
			FileName:  fpath,
			User:      username,
		}
		_, err = io.Copy(f, io.TeeReader(file, pr))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		db_entry := &database.File{
			Filepath:   fpath,
			Uploadtime: t,
		}
		err = database.NewFile(db_entry)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(map[string]string{
			"status": "uploaded",
			"path":   db_entry.Filepath,
			"time":   t.String(),
		})

	}

}
