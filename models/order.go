package models

import "time"

type Order struct {
	ID         string         `json:"id"`
	Reference  string       `json:"reference"`
	Customer   string       `json:"customer"`
	TotalPrice uint         `json:"-"`
	Date       time.Time    `json:"time"`
	Lines      []*OrderLine `json:"-"`
}
