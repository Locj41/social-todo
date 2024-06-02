package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"social-todo/modules/items/transport/gin"
)

func main() {
	dsn := "root:Qaz@040102@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	//Crud
	//post /v1/items (create new item)
	//get /v1/items	(get item details)
	//get /v1/items/:id (get item by id)
	//put /v1/items/:id (update item by id)
	//delete /v1/items/:id (delete item by id)

	v1 := router.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", ginitem.CreateItem(db))
			items.GET("", ginitem.GetAllItem(db))
			items.GET("/:id", ginitem.GetItemById(db))
			items.PUT("/:id", ginitem.UpdateItemById(db))
			items.DELETE("/:id", ginitem.DeleteItemById(db))
		}
	}

	router.GET("/", func(ctx *gin.Context) {
		ctx.File("templates/index.html")
	})

	router.GET("/profile", func(ctx *gin.Context) {
		ctx.File("templates/profile.html")
	})

	router.GET("/gallery", func(ctx *gin.Context) {
		ctx.File("templates/gallery.html")
	})

	router.GET("modal", func(ctx *gin.Context) {
		ctx.File("templates/modal.html")
	})

	router.Static("/static", "./static")

	router.Run()

}


