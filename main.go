package main

import (
	"net/http"
	"webapp/controller"
)

func main() {
	// 处理静态资源 static定位到views/static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	http.HandleFunc("/main", controller.IndexHandler)

	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/regist", controller.Regist)
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)
	http.HandleFunc("/updateOrAddBook", controller.UpdateOrAddBook)
	http.HandleFunc("/addBook2Cart", controller.AddBook2Cart)
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	http.HandleFunc("/deleteCart", controller.DeleteCart)
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)
	http.HandleFunc("/checkout", controller.Checkout)
	http.HandleFunc("/getOrders", controller.GetOrders)
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)
	http.HandleFunc("/getMyOrders", controller.GetMyOrders)
	http.HandleFunc("/sendOrder", controller.SendOrder)
	http.HandleFunc("/takeOrder", controller.TakeOrder)

	http.ListenAndServe(":8080", nil)
}
