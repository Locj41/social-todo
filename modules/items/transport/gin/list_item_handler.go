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

func GetAllItem(db *gorm.DB) func(c *gin.Context) {

	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBindQuery(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		paging.Process()

		var filter models.Filter
		if err := c.ShouldBindQuery(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		storage := storage.NewSQLStore(db)
		biz := business.NewGetAllItemBiz(storage)
		result, err := biz.GetAllItem(c.Request.Context(),&filter,&paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SuccessResponse(result, paging, filter))

	}
}
