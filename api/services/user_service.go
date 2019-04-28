package services

import (
	_ "io"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"Streaming-Media-Service/api/defs"
	"log"
	_ "Streaming-Media-Service/api/utils"
	"Streaming-Media-Service/api/dbops"
	"io/ioutil"
	"encoding/json"
)

//===========================================
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 首先从request里面读出create user相关的body
	res, _:=ioutil.ReadAll(r.Body)
	ubody:=&defs.UserCredential{}

	if err:=json.Unmarshal(res, ubody);err!=nil{
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err:=dbops.AddUser(ubody.Username, ubody.Pwd); err!=nil {
		sendErrorResponse(w, defs.ErrorDBError)
	}

	id:= GenerateNewSessionId(ubody.Username)
	su:=&defs.SignedUp{Success: true, SessionId: id}

	if resp, err:=json.Marshal(su);err!=nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _:=ioutil.ReadAll(r.Body)
	log.Printf("%s", res)
	ubody:=&defs.UserCredential{}
	if err:=json.Unmarshal(res, ubody); err!=nil {
		log.Printf("%s", err)
		//io.WriteString(w, "wrong")
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	uname := p.ByName("username")
	log.Printf("Login url name: %s", uname)
	log.Printf("Login body name: %s", ubody.Username)
	if uname!=ubody.Username {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}

	log.Printf("%s", ubody.Username)
	pwd, err:=dbops.GetUserCredential(ubody.Username)
	log.Printf("Login pwd: %s", pwd)
	log.Printf("Login body pwd: %s", ubody.Pwd)
	if err!=nil || len(pwd)==0 || pwd!=ubody.Pwd {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	id:= GenerateNewSessionId(ubody.Username)
	si := &defs.SignedIn{Success: true, SessionId: id}
	if resp, err:=json.Marshal(si); err!= nil{
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}

	//io.WriteString(w, uname)

}

func GetUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !ValidateUser(w, r) {
		log.Printf("Unauthorized user\n")
		return
	}

	uname:=p.ByName("username")
	u, err:=dbops.GetUser(uname)
	if err!=nil{
		log.Printf("Error in GetUserInfo: %s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	ui:=&defs.UserInfo{Id: u.Id}
	if resp, err:=json.Marshal(ui); err!= nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}
