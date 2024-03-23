package server

import "github.com/gin-gonic/gin"

func (server *Server) routes() {
	router := gin.Default()
	// User
	router.POST("/api/create/user", server.createUser)
	router.POST("/api/login/user", server.loginUser)
	router.POST("/api/update/password", server.updateuserPassword)

	// Learning data
	datarouter := router.Group("/api/data").Use(authMiddleware(server.token))
	datarouter.POST("/list-gender", server.listGender)
	datarouter.GET("/count-gender", server.countGender)
	datarouter.GET("/count-age", server.countAge)
	server.router = router
}
