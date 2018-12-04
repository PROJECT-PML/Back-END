package Handlers

import "net/http"

func GetLists(w http.ResponseWriter, req *http.Request) string{
	str := `API list:
prefix: /api

GET, POST /articles 
  GET, PUT /articles/{article_id}
    GET, POST /articles/{article_id}/comments
      GET, PUT /articles/{article_id}/comments/{comment_id}

GET, POST /users
  GET, PUT /users/{user_id}

GET, POST /tags
  GET, PUT /tags/{tag_id}

POST /auth
`
	return str
}

