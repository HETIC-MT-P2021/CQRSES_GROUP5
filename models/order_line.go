package models

type OrderLine struct {
	ID       uint   `json:"id"`
	Meal     string `json:"meal"`
	Quantity uint   `json:"quantity"`
	Price    uint   `json:"price"`
	IDOrder  uint   `json:"id_command"`
}
