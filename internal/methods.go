package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

var (
	ROOT_PATH string
	PORT      string = "8080"
	once      sync.Once
)

func GetRootPath() {
	once.Do(func() {
		ROOT_PATH = os.Getenv("ROOT_DIR")

		if ROOT_PATH == "" {
			panic("ROOT_DIR environment variable is not set")
		}

		if os.Getenv("PORT") != "" {
			PORT = os.Getenv("PORT")
		}
		fmt.Printf("Using ROOT_DIR: %s\n", ROOT_PATH)
		fmt.Printf("Using PORT: %s\n", PORT)
	})
}

func GetFile(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Query().Get("path")

	if filePath == "" {
		http.Error(w, "File path is required", http.StatusBadRequest)
		return
	}

	http.ServeFile(w, r, filePath)
}

func AddFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	path := r.FormValue("path")
	path = ROOT_PATH + path

	if err != nil {
		http.Error(w, "Failed to retrieve file from form data", http.StatusBadRequest)
		return
	}

	if file == nil || path == "" {
		http.Error(w, "File and path are required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		http.Error(w, "Failed to create directories", http.StatusInternalServerError)
		return
	}

	out, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to create file on server", http.StatusInternalServerError)
		return
	}
	defer out.Close()
	if _, err := io.Copy(out, file); err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("File uploaded successfully"))
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Query().Get("path")

	if filePath == "" {
		http.Error(w, "File path is required", http.StatusBadRequest)
		return
	}

	if err := os.Remove(filePath); err != nil {
		http.Error(w, "Failed to delete file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File deleted successfully"))
}
