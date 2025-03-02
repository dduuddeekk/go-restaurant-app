package model

type MenuType string

type MenuItem struct {
	Name      string   `json:"name"`
	OrderCode string   `json:"order_code"`
	Price     int64    `json:"price"`
	Type      MenuType `json:"type"`
}
