package dao

import (
	"fmt"
	"webapp/model"
	"webapp/utils"
)

func AddOrder(order *model.Order) error {
	sqlStr := "insert into orders(id,create_time,total_count,total_amount,state,user_id) values (?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, order.OrderID, order.CreateTime, order.TotalCount,
		order.TotalAmount, order.State, order.UserID)
	if err != nil {
		return err
	}
	return nil
}
func GetOrders() ([]*model.Order, error) {
	sqlStr := "select id,create_time,total_count,total_amount,state,user_id from orders"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		fmt.Print("GetOrders err:", err)
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount,
			&order.State, &order.UserID)
		orders = append(orders, order)
	}
	return orders, nil
}

func GetMyOrders(userID int64) ([]*model.Order, error) {
	sqlStr := "select id,create_time,total_count,total_amount,state,user_id from orders where user_id=?"
	rows, err := utils.Db.Query(sqlStr, userID)
	if err != nil {
		fmt.Print("GetOrders err:", err)
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount,
			&order.State, &order.UserID)
		orders = append(orders, order)
	}
	return orders, nil
}

func UpdateOrderState(orderID string, state int64) error {
	sqlStr := "update  orders set state =? where id=?"
	_, err := utils.Db.Exec(sqlStr, state, orderID)
	if err != nil {
		fmt.Print("UpdateOrderState err:", err)
		return err
	}
	return nil
}
