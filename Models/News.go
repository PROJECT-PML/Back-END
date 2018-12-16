package Models

import (
	"encoding/json"
	"strconv"
	"time"
	"../utils"
)

type NewsList struct {
	NewsID int    `json:"id"`
	Author    int    `json:"author_id"`
	Type int  	 `json:"type_id"`
	Title     string `json:"title"`
	UpdatedAt string `json:"updated_at"`
}

type News struct {
	NewsID int       `json:"id, omitempty"`
	Author    int       `json:"author_id"`
	Type int  	 `json:"type_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAllNews() []NewsList {
	db := &utils.DB{}
	var newss []NewsList
	var news NewsList
	var newssBytes map[string]string
	newssBytes = db.Scan("news")
	if len(newssBytes) == 0 {
		return []NewsList{}
	}
	for _, one := range newssBytes {
		err := json.Unmarshal([]byte(one), &news)
		newss = append(newss, news)
		if err != nil {
			panic(err)
		}
	}
	return newss
}

func CreateNews(news News) int {
	db := &utils.DB{}
	id := db.GenerateID("news")
	news.NewsID = id
	news.CreatedAt = time.Now()
	news.UpdatedAt = time.Now()
	buff, err := json.Marshal(news)
	if err != nil {
		panic("JSON parsing error")
	}
	db.Set("news", strconv.Itoa(id), string(buff))
	return id
}

func GetNewsByID(id int) *News {
	db := &utils.DB{}
	buff := db.Get("news", strconv.Itoa(id))
	if len(buff) == 0 {
		return nil
	}
	news := News{}
	err := json.Unmarshal(buff, &news)
	if err != nil {
		panic(err)
	}
	return &news
}

func UpdateNewsByID(id int, news News) bool {
	db := &utils.DB{}
	buff := db.Get("news", strconv.Itoa(id))
	if len(buff) == 0 {
		return false
	}
	buff, err := json.Marshal(news)
	if err != nil {
		panic("JSON parsing error")
	}
	db.Set("news", strconv.Itoa(id), string(buff))
	return true
}
