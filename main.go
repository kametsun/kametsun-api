package main

import (
	"kametsun-api/controllers"
	"kametsun-api/utils"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){
	db := utils.InitDataBase()
	defer db.Close()

	wishItemController := controllers.NewWishItemController(db)

	r := gin.Default()

	// CORSの設定
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "https://kametsun.vercel.app/"}
	r.Use(cors.New(config))

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	r.GET("/wishlist", wishItemController.GetWishItems)
	r.POST("/wishlist", wishItemController.Create)

	r.Run(":8080")
}
