package storage

import (
	"context"
	"social-todo/modules/items/models"
)

func (sql *sqlStore) CreateItem(ctx context.Context, data *models.TodoItemCreation) error {
	if err := sql.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
