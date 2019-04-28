package services

import (
	"io"
	"net/http"
	"encoding/json"
	"Streaming-Media-Service/api/defs"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrResponse) {
	w.WriteHeader(errResp.HttpSC)
	resStr, _:=json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}