package storage

import (
	"context"
	"social-todo/modules/items/models"
)

func (sql *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	if err := sql.db.
		Table(models.TodoItem{}.TableName()).
		Where(cond).
		Delete(nil).Error; err != nil {
			return err
	}
	return nil
}
