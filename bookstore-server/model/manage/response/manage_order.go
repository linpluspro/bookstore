package response

import "bookstore/model/common"

type BookStoreOrderDetailVO struct {
	OrderId                  int                     `json:"orderId"`
	OrderNo                  string                  `json:"orderNo"`
	TotalPrice               int                     `json:"totalPrice"`
	PayType                  int                     `json:"payType"`
	PayTypeString            string                  `json:"payTypeString"`
	OrderStatus              int                     `json:"orderStatus"`
	OrderStatusString        string                  `json:"orderStatusString"`
	CreateTime               common.JSONTime         `json:"createTime"`
	BookStoreOrderItemVOS    []BookStoreOrderItemVO  `json:"bookStoreOrderItemVOS"`
	BookStoreOrderAddressVOS BookStoreOrderAddressVO `json:"bookStoreOrderAddressVOS"`
}

type BookStoreOrderItemVO struct {
	BooksId       int    `json:"booksId"`
	BooksName     string `json:"booksName"`
	BooksCount    int    `json:"booksCount"`
	BooksCoverImg string `json:"booksCoverImg"`
	SellingPrice  int    `json:"sellingPrice"`
}

type BookStoreOrderAddressVO struct {
	AddressId     int    `json:"addressId"`
	UserId        int    `json:"userId"`
	UserName      string `json:"userName"`
	UserPhone     string `json:"userPhone"`
	DefaultFlag   int    `json:"defaultFlag"` // 0-不是 1-是
	ProvinceName  string `json:"provinceName"`
	CityName      string `json:"cityName"`
	RegionName    string `json:"regionName"`
	DetailAddress string `json:"detailAddress"`
}
