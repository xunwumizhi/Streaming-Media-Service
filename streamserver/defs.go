package main

import (
	"io"
	// "encoding/json"
	"net/http"
)

const (
	VIDEO_DIR = "./videos/"
	MAX_UPLOAD_SIZE = 1024*1024*50
)


func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}