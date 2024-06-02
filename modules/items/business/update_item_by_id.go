package business

import (
	"context"
	"social-todo/modules/items/models"
)

type UpdateItemStorage interface {
	GetItemStorage
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *models.TodoItemUpdate) error
}

type updateItemBiz struct {
	store UpdateItemStorage
}

func (biz *updateItemBiz) UpdateItemById(ctx context.Context, id int, dataUpdate *models.TodoItemUpdate) error {
	item, err := biz.store.GetItem(ctx, map[string]interface{}{"id":id})

	if err != nil {
		return err
	}	
	
	if item.Status != nil && *item.Status == models.ItemStatusDeleted {
		return models.ErrItemDeleted
	}

	if err := biz.store.UpdateItem(ctx, map[string]interface{}{"id":id}, dataUpdate); err != nil {
		return err
	}

	return nil 
	
}

func NewUpdateItemBiz(store UpdateItemStorage) *updateItemBiz{
	return &updateItemBiz{store: store}
}