package models

import "time"

//Order is the basic model for order ( which are sent in commands )
type Order struct {
	ID         uint         `json:"id"`
	Reference  string       `json:"reference"`
	Customer   string       `json:"customer"`
	TotalPrice uint         `json:"-"`
	Date       time.Time    `json:"time"`
	Lines      []*OrderLine `json:"-"`
}
