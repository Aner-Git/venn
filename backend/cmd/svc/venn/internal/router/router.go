package router

import (
	"github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/api"
	"github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func Setup(db *database.GormDatabase) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	//cors: to configure later...
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	neighborhoods := &api.NeighborhoodAPI{DB: db}
	//Get
	r.GET("/neighborhoods", parsePaginationDefault(), neighborhoods.GetClients)

	return r
}
