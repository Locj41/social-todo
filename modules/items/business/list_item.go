package business

import (
	"context"
	"social-todo/common"
	"social-todo/modules/items/models"
)

type GetAllItemStorage interface {
	ListItem(ctx context.Context,
		filter *models.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]models.TodoItem, error)
}

type getAllItemBiz struct {
	store GetAllItemStorage
}

func (biz *getAllItemBiz) GetAllItem(ctx context.Context,
	filter *models.Filter,
	paging *common.Paging,
) ([]models.TodoItem, error) {

	data, err := biz.store.ListItem(ctx, filter, paging)
	
	if err != nil {
		return nil, err
	}

	return data, nil

}

func NewGetAllItemBiz(store GetAllItemStorage) *getAllItemBiz {
	return &getAllItemBiz{store: store}
}
