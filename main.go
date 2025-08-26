package main

import (
	"fmt"
	"net/http"
	"tulucentre-fs/internal"
)

func main() {
	mux := http.NewServeMux()

	internal.GetRootPath()

	mux.HandleFunc("/api/fs", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Tulucentre File Server is running"))
	})
	mux.HandleFunc("/api/fs/getfile", internal.GetFile)
	mux.HandleFunc("/api/fs/addfile", internal.AddFile)
	// mux.HandleFunc("/api/delete", internal.DeleteFile)

	server := http.Server{
		Addr:    ":" + internal.PORT,
		Handler: mux,
	}

	fmt.Println("Starting server on :" + internal.PORT)
	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
