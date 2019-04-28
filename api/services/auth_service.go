package services

import (
	"net/http"
	"Streaming-Media-Service/api/defs"
)

//session============================================
var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

func ValidateUserSession(r *http.Request) bool {
	sid:=r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid)==0{
		return false
	}

	uname, ok := IsSessionExpired(sid)
	if ok {
		return false
	}

	//如果没有过期,那么就把用户名加入到HEADER_FIELD_UNAME里面
	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname:=r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname)==0{
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}

	return true
}
