package Handlers

import "net/http"

func GetLists(w http.ResponseWriter, req *http.Request) string{
	str := `API list:
prefix: /api

GET, POST /news 
GET /news/{type_id}
  GET, PUT /news/{news_id}
	
    GET, POST /news/{news_id}/comments
      GET, PUT /news/{news_id}/comments/{comment_id}

GET, POST /users
  GET, PUT /users/{user_id}


POST /auth
`
	return str
}

