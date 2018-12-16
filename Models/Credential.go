package Models

import (
	"encoding/json"
	myUtils "../utils" 
	"../defs"
	"net/http"
	"crypto/md5"
	"time"
	"strconv"
	"io"
	"fmt"
	"../utils"
)

var secretnum = []byte("sxkgjkcd abwye")


type AuthToken struct {
	// TokenID      int    `json:"id"`
	Token        string `json:"token"`
	AuthorizedID int    `json:"authorized_id"`
	ExpiredTime  string `json:"expired_time"`
}

func CreateToken(token AuthToken) error {
	db := &utils.DB{}
	buff, err := json.Marshal(token)
	if err != nil {
		return err
	}
	db.Set("token", token.Token, string(buff))
	return nil
}


func GetToken(tokenstr string) *AuthToken {
	db := &utils.DB{}
	buff := db.Get("token", tokenstr)
	if len(buff) == 0 {
		return nil
	}
	token := AuthToken{}
	if err := json.Unmarshal(buff, &token); err != nil {
		panic(err.Error())
	}
	return &token
}


func GenerateAuthToken(userID int) AuthToken {
	h := md5.New()
	expiredTime := time.Now().UnixNano()/1e6 + 1000*60*60*3 // expired time is 3 hours
	source := strconv.FormatInt(expiredTime, 10) + strconv.Itoa(userID)
	io.WriteString(h, source)
	token := fmt.Sprintf("%x", h.Sum(secretnum))
	authToken := AuthToken{
		// 0,
		token,
		userID,
		strconv.FormatInt(expiredTime, 10),
	}
	CreateToken(authToken)
	return authToken
}


func authenticateToken(token string) bool {
	real := GetToken(token)
	if real == nil {
		return false
	}
	if real.ExpiredTime < strconv.FormatInt((time.Now().UnixNano()/1e6), 10) {
		return false
	}
	return true
}

func GetCurrentUser(token string) *User {
	data := GetToken(token)
	user := GetUserByID(data.AuthorizedID)
	return user
}

func AuthenticationGuard(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error {
	header := req.Header
	token := header.Get("Authorization")
	if authenticateToken(token) {
		return next()
	} else {
		panic(utils.Exception{"Unauthorized", http.StatusUnauthorized})
	}
}


func ValidateUser(w http.ResponseWriter, r *http.Request) bool {

	header := r.Header
	token := header.Get("Authorization")
	if authenticateToken(token) == false {
		myUtils.SendErrorResponse(w, defs.ErrorNotAuthUser)
	}
	return true
}
