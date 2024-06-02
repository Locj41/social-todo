package business

import (
	"context"
	"social-todo/modules/items/models"
	"strings"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *models.TodoItemCreation) error
}

type createItemBiz struct {
	store CreateItemStorage
}

func (biz *createItemBiz) CreateNewItem(ctx context.Context, data *models.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title)

	if title == "" {
		return models.ErrTitileBlank
	}

	err := biz.store.CreateItem(ctx, data)
	if err != nil{
		return err
	}

	return nil
	
}

func NewCreateItemBiz(store CreateItemStorage) *createItemBiz{
	return &createItemBiz{store: store}
}