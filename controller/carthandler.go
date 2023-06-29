package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"webapp/dao"
	"webapp/model"
	"webapp/utils"
)

func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	flag, session := dao.IsLogin(r)
	if flag {
		bookID := r.PostFormValue("bookId")
		book, _ := dao.GetBookByID(bookID)
		userID := session.UserID
		cart, _ := dao.GetCartByUserID(userID)

		if cart != nil {
			// 当前客户有购物车
			cartItem, _ := dao.GetCartItemByBookIDAndCartID(bookID, cart.CartID)
			if cartItem != nil {
				//	购物车已有涂书，购物项+1
				cartItems := cart.CartItems
				for _, v := range cartItems {
					if v.Book.ID == cartItem.Book.ID {
						v.Count = v.Count + 1
						//	更新数据库中该数据项的图书的数量
						dao.UpdateBookCount(v)
					}
				}
			} else {
				//购物车的购物项没有该涂书，需要创建一个购物项并加入到数据库
				cartItem := &model.CartItem{
					Book:   book,
					Count:  1,
					CartID: cart.CartID,
				}
				cart.CartItems = append(cart.CartItems, cartItem)
				dao.AddCartItem(cartItem)
			}
			dao.UpdateCart(cart)
		} else {
			cartID := utils.CreateUUID()
			fmt.Printf("当前购物车id：%v\n", cartID)
			cart := &model.Cart{
				CartID: cartID,
				UserID: userID}
			var cartItems []*model.CartItem
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartID: cartID}
			cartItems = append(cartItems, cartItem)
			cart.CartItems = cartItems
			// 保持到数据库
			dao.AddCart(cart)
		}
		w.Write([]byte("您刚刚将" + book.Title + "添加到了购物车！"))
	} else {
		w.Write([]byte("请先登陆！"))
	}
}

func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	session.Cart = cart
	t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
	t.Execute(w, session)

}
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	cartID := r.FormValue("cartId")
	dao.DeleteCartByCartID(cartID)
	GetCartInfo(w, r)
}
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemId")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	cartItems := cart.CartItems
	for k, v := range cartItems {
		if v.CartItemID == iCartItemID {
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			cart.CartItems = cartItems
			dao.DeleteCartItemsByID(cartItemID)
		}
	}
	dao.UpdateCart(cart)
	GetCartInfo(w, r)
}
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemId")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	bookCount := r.FormValue("bookCount")
	iBookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	cartItems := cart.CartItems
	for _, v := range cartItems {
		if v.CartItemID == iCartItemID {
			v.Count = iBookCount
			//	 更新数据库中购物项
			dao.UpdateBookCount(v)
		}
	}
	dao.UpdateCart(cart)
	cart, _ = dao.GetCartByUserID(userID)
	totalCount := cart.TotalCount
	totalAmount := cart.TotalAmount
	var amount float64
	cIs := cart.CartItems
	for _, v := range cIs {
		if v.CartItemID == iCartItemID {
			amount = v.Amount
		}
	}
	data := model.Data{Amount: amount,
		TotalAmount: totalAmount,
		TotalCount:  totalCount}
	json, _ := json.Marshal(data)
	w.Write(json)
}
