package Handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"webPro/Models"
	"io/ioutil"
	"strconv"
	"webPro/utils"
	"webPro/defs"
	"io"
)

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}

func GetAllUsers(w http.ResponseWriter, req *http.Request, p httprouter.Params)  {
	res, err := json.Marshal(Models.GetAllUsers())
	if err != nil {
		panic(err)
	}
	utils.SendNormalResponse(w,string(res),http.StatusOK)
}

func CreateUser(w http.ResponseWriter, req *http.Request, p httprouter.Params)  {
	var user = Models.User{}
	buff, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(buff, &user)
	if err != nil {
		panic(err)
	}
	id := Models.CreateUser(user)
	utils.SendNormalResponse(w,`{"id":`+strconv.Itoa(id)+` }`,http.StatusOK)

}

func GetUserByID(w http.ResponseWriter, req *http.Request, p httprouter.Params)  {
	id := p.ByName("user_id")
	ID, err := strconv.Atoi(id)
	if err != nil{
		panic(err)
	}
	user := Models.GetUserByID(ID)
	if user == nil {
		utils.SendErrorResponse(w, defs.ErrorDBError)
	}
	data, err := json.Marshal(*user)
	if err != nil {
		utils.SendErrorResponse(w, defs.ErrorDBError)
	}
	utils.SendNormalResponse(w,string(data),http.StatusOK)
}

func UpdateUserByID(w http.ResponseWriter, req *http.Request, p httprouter.Params)  {
	var user = Models.User{}
	buff, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	id := p.ByName("user_id")
	ID, err := strconv.Atoi(id)
	if err != nil{
		panic(err)
	}
	user.UserID = ID
	err = json.Unmarshal(buff, &user)
	if err != nil {
		panic(err)
	}
	isUpdated := Models.UpdateUserByID(ID, user)
	if !isUpdated {
		ID = Models.CreateUser(user)
		utils.SendNormalResponse(w,`{"id": "`+id+`"}`,http.StatusCreated)

	}
	utils.SendNormalResponse(w,`{}`,http.StatusOK)

}

