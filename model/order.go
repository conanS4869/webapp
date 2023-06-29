package model

type Order struct {
	OrderID     string
	CreateTime  string
	TotalCount  int64
	TotalAmount float64
	State       int64
	UserID      int64
}

// NoSend 未发货
func (order *Order) NoSend() bool {
	return order.State == 0
}

// SendComplete 已发货
func (order *Order) SendComplete() bool {
	return order.State == 1
}

// Complete 交易完成
func (order *Order) Complete() bool {
	return order.State == 2
}
