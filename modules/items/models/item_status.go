package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
)


var allItemStatus = [3]string{"doing", "done", "deleted"}

type ItemStatus int

// convert struct to json data
func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

//convert json data to struct

func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")
	itemValue, err := parseStringtoItemStatus(str)
	if err != nil {
		return err
	}
	*item = itemValue
	return nil
}

// convert value of struct to database value
func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}
	return item.String(), nil
}

//scan value retrieve from database to itemstatus(int)
//value : ['doing', 'done', 'deleted'] --> []byte

// convert itemstatus to string
func (item ItemStatus) String() string {
	return allItemStatus[item]
}

func parseStringtoItemStatus(s string) (ItemStatus, error) {
	for i := range allItemStatus {
		if allItemStatus[i] == s {
			//convert i to itemstatus although they just one :)
			return ItemStatus(i), nil
		}
	}

	return ItemStatus(0), fmt.Errorf("INVALID STRING: %s", s)
}

func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return fmt.Errorf("FAILED TO SCAN VALUE: %s", value)
	}

	strValue := string(bytes)
	v, err := parseStringtoItemStatus(strValue)

	if err != nil {
		return err
	}
	*item = v
	return nil
}

const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

