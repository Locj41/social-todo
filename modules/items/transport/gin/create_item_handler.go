package ginitem

import (
	"net/http"
	"social-todo/common"
	"social-todo/modules/items/business"
	"social-todo/modules/items/models"
	"social-todo/modules/items/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(c *gin.Context) {

	return func(c *gin.Context) {
		var data models.TodoItemCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		storage := storage.NewSQLStore(db)
		biz := business.NewCreateItemBiz(storage)
		if err := biz.CreateNewItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
