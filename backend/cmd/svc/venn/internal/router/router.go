package router

import (
	"github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/api"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func Setup( /*pass db */ ) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	neighborhoods := &api.NeighborhoodAPI{}
	//Get
	r.GET("/neighborhoods", neighborhoods.GetClients)

	return r
}
