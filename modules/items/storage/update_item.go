package storage

import (
	"context"
	"social-todo/modules/items/models"
)

func (sql *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *models.TodoItemUpdate) error {
	err := sql.db.Where(cond).Updates(dataUpdate).Error
	if err != nil {
		return err
	}
	return nil
}
