package storage

import (
	"context"
	"social-todo/modules/items/models"
)

func (sql *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*models.TodoItem, error) {
	var data models.TodoItem

	if err := sql.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}
	return &data,nil
}
