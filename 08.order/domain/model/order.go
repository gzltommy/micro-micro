package model

import "time"

type Order struct {
	ID          int64         `gorm:"primary_key;not_null;auto_increment",json:"id"`
	OrderCode   string        `gorm:"unique_index;not_null",json:"order_code"`
	PayStatus   int32         `json:"pay_status"`  // 支付状态
	ShipStatus  int32         `json:"ship_status"` // 发货状态
	Price       float64       `json:"price"`
	OrderDetail []OrderDetail `gorm:"ForeignKey:OrderID" json:"order_detail"`
	CreateAt    time.Time
	UpdateAt    time.Time
}
