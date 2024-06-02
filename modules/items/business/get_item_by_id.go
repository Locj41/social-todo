package business

import (
	"context"
	"social-todo/modules/items/models"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*models.TodoItem,error)
}

type getItemBiz struct {
	store GetItemStorage
}

func (biz *getItemBiz) GetItemById(ctx context.Context, id int) (*models.TodoItem,error) {
	item, err := biz.store.GetItem(ctx, map[string]interface{}{"id":id})
	if err != nil {
		return nil, err
	}	
	
	return item, nil 
	
}

func NewGetItemBiz(store GetItemStorage) *getItemBiz{
	return &getItemBiz{store: store}
}