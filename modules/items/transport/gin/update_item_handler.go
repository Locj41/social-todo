package ginitem

import (
	"net/http"
	"social-todo/common"
	"social-todo/modules/items/business"
	"social-todo/modules/items/models"
	"social-todo/modules/items/storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateItemById(db *gorm.DB) func(c *gin.Context) {

	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var data models.TodoItemUpdate
		err_parse := c.ShouldBind(&data)
		if err_parse != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		storage := storage.NewSQLStore(db)
		biz := business.NewUpdateItemBiz(storage)
		if err := biz.UpdateItemById(c.Request.Context(), id, &data);err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}