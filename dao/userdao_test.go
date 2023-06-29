package dao

import (
	"fmt"
	"testing"
	"webapp/model"
)

func TestCartItem(t *testing.T) {
	t.Run("testGetCartItemByCartID", testGetCartItemsByCartID)
	t.Run("testGetCartByUserID", testGetCartByUserID)
}
func testAddCart(t *testing.T) {
	book := &model.Book{
		ID:    1,
		Price: 10.00,
	}
	book2 := &model.Book{
		ID:    2,
		Price: 18.88,
	}
	var cartItems []*model.CartItem
	cartItem := &model.CartItem{
		Book:   book,
		Count:  10,
		CartID: "1",
	}
	cartItems = append(cartItems, cartItem)
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  10,
		CartID: "1",
	}
	cartItems = append(cartItems, cartItem2)

	cart := &model.Cart{
		UserID:    1,
		CartID:    "1",
		CartItems: cartItems}
	AddCart(cart)
}

func testGetCartItemByBookID(t *testing.T) {
	cartItem, _ := GetCartItemByBookID("1")
	fmt.Println("TestGetCartItemByBookID", cartItem)
}
func testGetCartItemsByCartID(t *testing.T) {
	cartItems, _ := GetCartItemsByCartID("693c78d6-c20e-48ee-792e-a58419de294d")
	for i, cartItem := range cartItems {
		fmt.Printf("第%v个购物项是: %v\n", i, cartItem)
	}
}
func testGetCartByUserID(t *testing.T) {
	cart, _ := GetCartByUserID(1)
	fmt.Println("id为1的用户的购物车为:", cart)
}
