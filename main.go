package main

import (
	"fmt"
	"net/http"
	"tulucentre-fs/internal"
)

func main() {
	mux := http.NewServeMux()

	internal.GetRootPath()

	mux.HandleFunc("/api/getfile", internal.GetFile)
	mux.HandleFunc("/api/addfile", internal.AddFile)
	mux.HandleFunc("/api/delete", internal.DeleteFile)

	server := http.Server{
		Addr:    ":" + internal.PORT,
		Handler: mux,
	}

	fmt.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
