package mall

import (
	"errors"
	"time"

	"github.com/jinzhu/copier"

	"bookstore/global"
	"bookstore/model/common"
	"bookstore/model/mall"
	mallReq "bookstore/model/mall/request"
	mallRes "bookstore/model/mall/response"
	"bookstore/model/manage"
)

type MallShopCartService struct {
}

// GetMyShoppingCartItems 不分页
func (m *MallShopCartService) GetMyShoppingCartItems(token string) (err error, cartItems []mallRes.CartItemResponse) {
	var userToken mall.MallUserToken
	var shopCartItems []mall.MallShoppingCartItem
	var booksInfos []manage.MallBooksInfo
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), cartItems
	}
	global.GVA_DB.Where("user_id=? and is_deleted = 0", userToken.UserId).Find(&shopCartItems)
	var booksIds []int
	for _, shopcartItem := range shopCartItems {
		booksIds = append(booksIds, shopcartItem.BooksId)
	}
	global.GVA_DB.Where("books_id in ?", booksIds).Find(&booksInfos)
	booksMap := make(map[int]manage.MallBooksInfo)
	for _, booksInfo := range booksInfos {
		booksMap[*booksInfo.BooksId] = booksInfo
	}
	for _, v := range shopCartItems {
		var cartItem mallRes.CartItemResponse
		copier.Copy(&cartItem, &v)
		if _, ok := booksMap[v.BooksId]; ok {
			booksInfo := booksMap[v.BooksId]
			cartItem.BooksName = booksInfo.BooksName
			cartItem.BooksCoverImg = booksInfo.BooksCoverImg
			cartItem.SellingPrice = *booksInfo.SellingPrice
		}
		cartItems = append(cartItems, cartItem)
	}

	return
}

func (m *MallShopCartService) SaveMallCartItem(token string, req mallReq.SaveCartItemParam) (err error) {
	if req.BooksCount < 1 {
		return errors.New("商品数量不能小于 1 ！")

	}
	if req.BooksCount > 5 {
		return errors.New("超出单个商品的最大购买数量！")
	}
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}
	var shopCartItems []mall.MallShoppingCartItem
	// 是否已存在商品
	err = global.GVA_DB.Where("user_id = ? and books_id = ? and is_deleted = 0", userToken.UserId, req.BooksId).Find(&shopCartItems).Error
	if err != nil {
		return errors.New("已存在！无需重复添加！")
	}
	err = global.GVA_DB.Where("books_id = ? ", req.BooksId).First(&manage.MallBooksInfo{}).Error
	if err != nil {
		return errors.New(" 商品为空")
	}
	var total int64
	global.GVA_DB.Where("user_id =?  and is_deleted = 0", userToken.UserId).Count(&total)
	if total > 20 {
		return errors.New("超出购物车最大容量！")
	}
	var shopCartItem mall.MallShoppingCartItem
	if err = copier.Copy(&shopCartItem, &req); err != nil {
		return err
	}
	shopCartItem.UserId = userToken.UserId
	shopCartItem.CreateTime = common.JSONTime{Time: time.Now()}
	shopCartItem.UpdateTime = common.JSONTime{Time: time.Now()}
	err = global.GVA_DB.Save(&shopCartItem).Error
	return
}

func (m *MallShopCartService) UpdateMallCartItem(token string, req mallReq.UpdateCartItemParam) (err error) {
	//超出单个商品的最大数量
	if req.BooksCount > 5 {
		return errors.New("超出单个商品的最大购买数量！")
	}
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}
	var shopCartItem mall.MallShoppingCartItem
	if err = global.GVA_DB.Where("cart_item_id=? and is_deleted = 0", req.CartItemId).First(&shopCartItem).Error; err != nil {
		return errors.New("未查询到记录！")
	}
	if shopCartItem.UserId != userToken.UserId {
		return errors.New("禁止该操作！")
	}
	shopCartItem.BooksCount = req.BooksCount
	shopCartItem.UpdateTime = common.JSONTime{time.Now()}
	err = global.GVA_DB.Save(&shopCartItem).Error
	return
}

func (m *MallShopCartService) DeleteMallCartItem(token string, id int) (err error) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}
	var shopCartItem mall.MallShoppingCartItem
	err = global.GVA_DB.Where("cart_item_id = ? and is_deleted = 0", id).First(&shopCartItem).Error
	if err != nil {
		return
	}
	if userToken.UserId != shopCartItem.UserId {
		return errors.New("禁止该操作！")
	}
	err = global.GVA_DB.Where("cart_item_id = ? and is_deleted = 0", id).UpdateColumns(&mall.MallShoppingCartItem{IsDeleted: 1}).Error
	return
}

func (m *MallShopCartService) GetCartItemsForSettle(token string, cartItemIds []int) (err error, cartItemRes []mallRes.CartItemResponse) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), cartItemRes
	}
	var shopCartItems []mall.MallShoppingCartItem
	err = global.GVA_DB.Where("cart_item_id in (?) and user_id = ? and is_deleted = 0", cartItemIds, userToken.UserId).Find(&shopCartItems).Error
	if err != nil {
		return
	}
	_, cartItemRes = getMallShoppingCartItemVOS(shopCartItems)
	//购物车算价
	priceTotal := 0
	for _, cartItem := range cartItemRes {
		priceTotal = priceTotal + cartItem.BooksCount*cartItem.SellingPrice
	}
	return
}

// 购物车数据转换
func getMallShoppingCartItemVOS(cartItems []mall.MallShoppingCartItem) (err error, cartItemsRes []mallRes.CartItemResponse) {
	var booksIds []int
	for _, cartItem := range cartItems {
		booksIds = append(booksIds, cartItem.BooksId)
	}
	var bookStoreBooks []manage.MallBooksInfo
	err = global.GVA_DB.Where("books_id in ?", booksIds).Find(&bookStoreBooks).Error
	if err != nil {
		return
	}

	bookStoreBooksMap := make(map[int]manage.MallBooksInfo)
	for _, booksInfo := range bookStoreBooks {
		bookStoreBooksMap[*booksInfo.BooksId] = booksInfo
	}
	for _, cartItem := range cartItems {
		var cartItemRes mallRes.CartItemResponse
		copier.Copy(&cartItemRes, &cartItem)
		// 是否包含key
		if _, ok := bookStoreBooksMap[cartItemRes.BooksId]; ok {
			bookStoreBooksTemp := bookStoreBooksMap[cartItemRes.BooksId]
			cartItemRes.BooksCoverImg = bookStoreBooksTemp.BooksCoverImg
			booksName := bookStoreBooksTemp.BooksName
			cartItemRes.BooksName = booksName
			cartItemRes.SellingPrice = *bookStoreBooksTemp.SellingPrice
			cartItemsRes = append(cartItemsRes, cartItemRes)
		}
	}
	return
}
