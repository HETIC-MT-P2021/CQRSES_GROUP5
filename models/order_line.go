package models

//OrderLine is the basic model for order lines ( which are sent in commands )
type OrderLine struct {
	ID       uint   `json:"id"`
	Meal     string `json:"meal"`
	Quantity uint   `json:"quantity"`
	Price    uint   `json:"price"`
	OrderID  uint   `json:"order_id"`
}
