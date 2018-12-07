package Models

import (
	"encoding/json"
	"strconv"
	"time"
	"../utils"
)

type UserList struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
}

type User struct {
	UserID    int       `json:"id, omitempty"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func GetAllUsers() []UserList {
	db := &utils.DB{}
	var Users []UserList
	var user UserList
	var UsersBytes map[string]string
	UsersBytes = db.Scan("user")
	if len(UsersBytes) == 0 {
		return []UserList{}
	}
	for _, one := range UsersBytes {
		err := json.Unmarshal([]byte(one), &user)
		Users = append(Users, user)
		if err != nil {
			panic(err)
		}
	}
	return Users
}

func CreateUser(user User) int {
	db := &utils.DB{}
	id := db.GenerateID("user")
	user.UserID = id
	user.CreatedAt = time.Now()
	buff, err := json.Marshal(user)
	if err != nil {
		panic("JSON parsing error")
	}
	db.Set("user", strconv.Itoa(id), string(buff))
	return id
}

func GetUserByID(id int) *User {
	db := &utils.DB{}
	buff := db.Get("user", strconv.Itoa(id))
	if len(buff) == 0 {
		return nil
	}
	user := User{}
	err := json.Unmarshal(buff, &user)
	if err != nil {
		panic(err)
	}
	return &user
}

func UpdateUserByID(id int, user User) bool {
	db := &utils.DB{}
	buff := db.Get("user", strconv.Itoa(id))
	if len(buff) == 0 {
		return false
	}
	buff, err := json.Marshal(user)
	if err != nil {
		panic("JSON parsing error")
	}
	db.Set("user", strconv.Itoa(id), string(buff))
	return true
}

func DeleteUserByID(id int, user User) bool{
	db := &utils.DB{}
	buff := db.Get("user", strconv.Itoa(id))
	if len(buff) == 0 {
		return false
	}
	buff, err := json.Marshal(user)
	if err != nil {
		panic("JSON parsing error")
	}
	db.Delete("user")
	return true
}