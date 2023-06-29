package controller

import (
	"html/template"
	"net/http"
	"strconv"
	"webapp/dao"
	"webapp/model"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	pageNo := r.PostFormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageBooks(pageNo)
	flag, session := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.UserName = session.UserName
	}
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)
}

func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageBooks(pageNo)
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, page)
}
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}
	// 判断是否登录
	flag, session := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.UserName = session.UserName
	}
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	BookID := r.FormValue("bookId")
	dao.DeleteBook(BookID)
	GetPageBooks(w, r)
}

func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	BookID := r.FormValue("bookId")
	book, _ := dao.GetBookByID(BookID)
	if book.ID > 0 {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, book)
	} else {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, "")
	}

}
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	bookID := r.PostFormValue("bookId")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)
	ibookID, _ := strconv.ParseInt(bookID, 10, 0)
	book := &model.Book{
		ID:      ibookID,
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   iSales,
		Stock:   iStock,
		ImgPath: "/static/img/default.jpg"}
	if book.ID > 0 {
		dao.UpdateBook(book)
	} else {
		dao.AddBook(book)
	}
	GetPageBooks(w, r)
}
