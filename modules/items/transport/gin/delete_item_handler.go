package ginitem

import (
	"net/http"
	"social-todo/common"
	"social-todo/modules/items/business"
	"social-todo/modules/items/storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteItemById(db *gorm.DB) func(c *gin.Context) {

	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		storage := storage.NewSQLStore(db)
		biz := business.NewDeleteItemBiz(storage)

		if err := biz.DeleteItemById(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
