package business

import (
	"context"
	"social-todo/modules/items/models"
)

type DeleteItemStorage interface {
	GetItemStorage
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}

type deleteItemBiz struct {
	store DeleteItemStorage
}


func (biz *deleteItemBiz) DeleteItemById(ctx context.Context, id int) error {
	item, err := biz.store.GetItem(ctx, map[string]interface{}{"id":id})

	if err != nil {
		return err
	}	
	
	if item.Status != nil && *item.Status == models.ItemStatusDeleted {
		return models.ErrItemDeleted
	}

	if err := biz.store.DeleteItem(ctx, map[string]interface{}{"id":id}); err != nil {
		return err
	}

	return nil 
	
}

func NewDeleteItemBiz(store DeleteItemStorage) *deleteItemBiz{
	return &deleteItemBiz{store: store}
}