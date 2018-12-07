package Handlers

import (
	"net/http"
	"io/ioutil"
	"../Models"
	"encoding/json"
	"strconv"
	"github.com/julienschmidt/httprouter"
	"../utils"
)

func CreateComment(w http.ResponseWriter, req *http.Request, p httprouter.Params)  {

	newsID := p.ByName("news_id")
	author := Models.GetCurrentUser(req.Header.Get("Authorization"))
	buff, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	comment := Models.Comment{}
	err = json.Unmarshal(buff, &comment)
	if err != nil {
		panic(err)
	}
	comment.Creator = author.UserID
	NewsID, err := strconv.Atoi(newsID)
	if err != nil{
		panic(err)
	}
	comment.NewsID =  NewsID
	id := Models.CreateComment(&comment)
	utils.SendNormalResponse(w,`{"id":`+strconv.Itoa(id)+` }`,http.StatusOK)
}

func GetAllCommentsbyID(w http.ResponseWriter, req *http.Request, p httprouter.Params)  {

	newsID := p.ByName("news_id")
	NewsID, err := strconv.Atoi(newsID)
	if err != nil{
		panic(err)
	}
	comments := Models.GetAllCommentsByNewsID(NewsID)
	var buff []interface{}
	for _, one := range comments {
		buff = append(buff, one)
	}
	data, err := json.Marshal(buff)
	if err != nil {
		panic(err)
	}
	utils.SendNormalResponse(w,string(data),http.StatusOK)
}

