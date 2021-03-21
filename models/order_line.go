package models

//OrderLine is the basic model for order lines ( which are sent in commands )
type OrderLine struct {
	ID       string `json:"id"`
	Meal     string `json:"meal"`
	Quantity uint   `json:"quantity"`
	Price    uint   `json:"price"`
	OrderID  string `json:"order_id"`
}
