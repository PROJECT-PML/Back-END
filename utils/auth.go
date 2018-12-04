package utils

import (
	"net/http"
	"crypto/md5"
	"time"
	"strconv"
	"io"
	"fmt"
	"webPro/defs"
	"webPro/Models"
	"github.com/leiysky/go-cloud-service/utils"
)

var checksum = []byte("leiysky-is-a-handsome-boy")

func GenerateAuthToken(userID int) Models.AuthToken {
	h := md5.New()
	expiredTime := time.Now().UnixNano()/1e6 + 1000*60*60*3 // expired time is 3 hours
	source := strconv.FormatInt(expiredTime, 10) + strconv.Itoa(userID)
	io.WriteString(h, source)
	token := fmt.Sprintf("%x", h.Sum(checksum))
	authToken := Models.AuthToken{
		// 0,
		token,
		userID,
		strconv.FormatInt(expiredTime, 10),
	}
	Models.CreateToken(authToken)
	return authToken
}

func authenticateToken(token string) bool {
	real := Models.GetToken(token)
	if real == nil {
		return false
	}
	if real.ExpiredTime < strconv.FormatInt((time.Now().UnixNano()/1e6), 10) {
		return false
	}
	return true
}

func GetCurrentUser(token string) *Models.User {
	data := Models.GetToken(token)
	user := Models.GetUserByID(data.AuthorizedID)
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
		SendErrorResponse(w, defs.ErrorNotAuthUser)
	}
	return true
}
