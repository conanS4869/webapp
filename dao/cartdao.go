package dao

import (
	"webapp/model"
	"webapp/utils"
)

func UpdateCart(cart *model.Cart) error {
	sqlStr := "update carts set total_count=?,total_amount=? where id=?"
	_, err := utils.Db.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return err
	}
	return nil
}

func AddCart(cart *model.Cart) error {
	sqlStr := "insert into carts(id,total_count,total_amount,user_id)values (?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		return err
	}
	cartItems := cart.CartItems
	for _, cartItem := range cartItems {
		AddCartItem(cartItem)
	}
	return nil
}
func GetCartByUserID(userID int64) (*model.Cart, error) {
	sqlStr := "select id, total_count, total_amount, user_id from carts where user_id=?"
	row := utils.Db.QueryRow(sqlStr, userID)
	cart := &model.Cart{}
	err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil, err
	}
	cartItems, _ := GetCartItemsByCartID(cart.CartID)
	cart.CartItems = cartItems
	return cart, err
}
func DeleteCartByCartID(cartID string) error {
	err := DeleteCartItemsByCartID(cartID)
	if err != nil {
		return err
	}
	sqlStr := "delete from carts where id =?"
	_, err = utils.Db.Exec(sqlStr, cartID)
	if err != nil {
		return err
	}
	return nil
}
