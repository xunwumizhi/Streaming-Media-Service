package main

import (
	"io"
	"net/http"
	"os"
	"time"
	"html/template"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
	"log"
	// "strings"
)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _:=template.ParseFiles("./videos/upload.html")
	t.Execute(w, nil)
}

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")

	// comma := strings.Index(vid, ".")
	// codefmt := vid[comma+1:]

	vl := VIDEO_DIR + vid

	log.Printf("file name : %s", vl)
	// log.Printf("file name : %s", codefmt)

	video, err := os.Open(vl)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "error of open video: "+err.Error())
		return
	}
	w.Header().Set("Content-Type", "video/mp4")
	// w.Header().Set("Content-Type", "video/"+codefmt)
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err:=r.ParseMultipartForm(MAX_UPLOAD_SIZE); err!=nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too large")
		return
	}

	//此处一定要这么写，否则读不出文件
	file, _, err := r.FormFile("file")
	if err!=nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
	}

	data, err:=ioutil.ReadAll(file)
	if err!=nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
	}

	fn:=p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR + fn, data, 06666)
	if err!=nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully.")
}
