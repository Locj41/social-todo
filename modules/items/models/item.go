package models

import (
	"errors"
	"social-todo/common"
)

var (
	ErrTitileBlank = errors.New("title is cannot blank")
	ErrItemDeleted = errors.New("item deleted")
)


type TodoItem struct {
	common.SQLModel
	Title       string      `json:"title" gorm:"columnn:title"`
	Description string      `json:"description" gorm:"columnn:description"`
	Status      *ItemStatus `json:"status" gorm:"columnn:status"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

type TodoItemCreation struct {
	Id          int         `json:"_" gorm:"columnn:id"`
	Title       string      `json:"title" gorm:"column:title"`
	Description *string     `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}

func (TodoItemCreation) TableName() string {
	return TodoItem{}.TableName()
}

type TodoItemUpdate struct {
	Title       string      `json:"title" gorm:"columnn:title"`
	Description *string     `json:"description" gorm:"columnn:description"`
	Status      *ItemStatus `json:"status" gorm:"columnn:status"`
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}