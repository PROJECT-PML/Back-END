package Handlers

import (
	"net/http"
	"webPro/Models"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"github.com/julienschmidt/httprouter"
	"time"
	"webPro/defs"
	"webPro/utils"
)

func GetAllNews(w http.ResponseWriter, req *http.Request) {
	res, err := json.Marshal(Models.GetAllNews())
	if err != nil {
		panic(err)
	}
	 utils.SendNormalResponse(w, string(res), http.StatusOK)
}

func CreateNew(w http.ResponseWriter, req *http.Request)  {
	var news = Models.News{}
	verifier := utils.GetCurrentUser(req.Header.Get("Authorization"))
	buff, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(buff, &news)
	if err != nil {
		panic(err)
	}
	news.Verifier = verifier.UserID
	id := Models.CreateNews(news)
	utils.SendNormalResponse(w,`{
  	"id": `+strconv.Itoa(id)+`
		}`,http.StatusOK)

}

func GetNewsByID(w http.ResponseWriter, req *http.Request,p httprouter.Params)  {
	id := p.ByName("user_id")
	ID, err := strconv.Atoi(id)
	if err != nil{
		panic(err)
	}
	news := Models.GetNewsByID(ID)
	if news == nil {
		utils.SendErrorResponse(w, defs.ErrorInternalFaults)
	}
	author := Models.GetUserByID(news.Author)
	result := make(map[string]interface{})
	result["id"] = news.NewsID
	result["title"] = news.Title
	result["content"] = news.Content
	result["author_id"] = news.Author
	result["author"] = author.Username
	result["created_at"] = news.CreatedAt
	result["updated_at"] = news.UpdatedAt
	data, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	utils.SendNormalResponse(w,string(data),http.StatusOK)

}

func UpdateNewByID(w http.ResponseWriter, req *http.Request,p httprouter.Params){
	var news = Models.News{}
	buff, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	id := p.ByName("user_id")
	ID, err := strconv.Atoi(id)
	if err != nil{
		panic(err)
	}
	news.NewsID = ID
	news.UpdatedAt = time.Now()
	err = json.Unmarshal(buff, &news)
	if err != nil {
		panic(err)
	}
	isUpdated := Models.UpdateNewsByID(ID, news)
	if !isUpdated {
		ID = Models.CreateNews(news)
		utils.SendNormalResponse(w,`{"id": "`+strconv.Itoa(ID)+`"}`,http.StatusCreated)
	}
	utils.SendNormalResponse(w,"{}",http.StatusOK)
}

