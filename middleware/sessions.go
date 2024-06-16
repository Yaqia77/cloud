package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var store sessions.Store

func InitSessionStore() {
	var err error
	store, err = redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		panic(err)
	}
}

func SessionMiddleware() gin.HandlerFunc {
	if store == nil {
		panic("Session store not initialized")
	}
	return sessions.Sessions("mysession", store)
}
