package model

type Session struct {
	SessionID string
	UserName  string
	UserID    int64
	Cart      *Cart // 购物车
	Order     *Order
	Orders    []*Order
}
