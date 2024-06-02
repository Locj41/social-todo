package storage

import (
	"context"
	"social-todo/common"
	"social-todo/modules/items/models"
)

func (s *sqlStore) ListItem(ctx context.Context,
	filter *models.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]models.TodoItem, error) {

	var result []models.TodoItem
	//where status != deleted
	db := s.db.Where("status <> ?", "deleted")

	//if wnat to filter more (eg: status = done or doing)
	//result in many filter and condition will store in db var
	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}
	
	if err := db.Table(models.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	if err := db.
		Order("id asc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil

}
