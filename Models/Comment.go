package Models


import (
"encoding/json"
"strconv"
"time"

"../utils"
)

type Comment struct {
	CommentID int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	NewsID int       `json:"News_id"`
	Creator   int       `json:"creator"`
}

func CreateComment(comment *Comment) int {
	db := &utils.DB{}
	id := db.GenerateID("comment")
	comment.CommentID = id
	comment.CreatedAt = time.Now()
	buff, err := json.Marshal(comment)
	if err != nil {
		panic(err)
	}
	db.Set("comment", strconv.Itoa(id), string(buff))
	return id
}

func GetAllCommentsByNewsID(NewsID int) []Comment {
	db := &utils.DB{}
	comments := db.Scan("comment")
	var result []Comment
	for _, v := range comments {
		tmp := Comment{}
		err := json.Unmarshal([]byte(v), &tmp)
		if err != nil {
			panic(err)
		}
		if tmp.NewsID == NewsID {
			result = append(result, tmp)
		}
	}
	return result
}
