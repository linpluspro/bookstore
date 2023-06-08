package mall

import (
	"errors"
	"time"

	"github.com/jinzhu/copier"

	"bookstore/global"
	"bookstore/model/common"
	"bookstore/model/common/enum"
	"bookstore/model/mall"
	mallRes "bookstore/model/mall/response"
	"bookstore/model/manage"
	manageReq "bookstore/model/manage/request"
	"bookstore/utils"
)

type MallOrderService struct {
}

// SaveOrder 保存订单
func (m *MallOrderService) SaveOrder(token string, userAddress mall.MallUserAddress, myShoppingCartItems []mallRes.CartItemResponse) (err error, orderNo string) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), orderNo
	}
	var itemIdList []int
	var booksIds []int
	for _, cartItem := range myShoppingCartItems {
		itemIdList = append(itemIdList, cartItem.CartItemId)
		booksIds = append(booksIds, cartItem.BooksId)
	}
	var bookStoreBooks []manage.MallBooksInfo
	global.GVA_DB.Where("books_id in ? ", booksIds).Find(&bookStoreBooks)

	//检查是否包含已下架商品
	for _, mallBooks := range bookStoreBooks {
		if mallBooks.BooksSellStatus != enum.BOOKS_UNDER.Code() {
			return errors.New("已下架，无法生成订单"), orderNo
		}
	}
	bookStoreBooksMap := make(map[int]manage.MallBooksInfo)
	for _, mallBooks := range bookStoreBooks {
		bookStoreBooksMap[*mallBooks.BooksId] = mallBooks
	}

	//判断商品库存
	for _, shoppingCartItemVO := range myShoppingCartItems {
		//查出的商品中不存在购物车中的这条关联商品数据，直接返回错误提醒
		if _, ok := bookStoreBooksMap[shoppingCartItemVO.BooksId]; !ok {
			return errors.New("购物车数据异常！"), orderNo
		}
		if shoppingCartItemVO.BooksCount > *bookStoreBooksMap[shoppingCartItemVO.BooksId].StockNum {
			return errors.New("库存不足！"), orderNo
		}
	}

	//删除购物项
	if len(itemIdList) > 0 && len(booksIds) > 0 {
		if err = global.GVA_DB.Where("cart_item_id in ?", itemIdList).Updates(mall.MallShoppingCartItem{IsDeleted: 1}).Error; err == nil {
			var stockNumDTOS []manageReq.StockNumDTO
			copier.Copy(&stockNumDTOS, &myShoppingCartItems)
			for _, stockNumDTO := range stockNumDTOS {
				var booksInfo manage.MallBooksInfo
				global.GVA_DB.Where("books_id =?", stockNumDTO.BooksId).First(&booksInfo)
				tmp := *booksInfo.StockNum - stockNumDTO.BooksCount
				if err = global.GVA_DB.Where("books_id =? and stock_num>= ? and books_sell_status = 0",
					stockNumDTO.BooksId, stockNumDTO.BooksCount).Updates(
					manage.MallBooksInfo{StockNum: &tmp}).Error; err != nil {
					return errors.New("库存不足！"), orderNo
				}
			}
			//生成订单号
			orderNo = utils.GenOrderNo()
			priceTotal := 0
			//保存订单
			var bookStoreOrder manage.MallOrder
			bookStoreOrder.OrderNo = orderNo
			bookStoreOrder.UserId = userToken.UserId
			//总价
			for _, bookStoreShoppingCartItemVO := range myShoppingCartItems {
				priceTotal = priceTotal + bookStoreShoppingCartItemVO.BooksCount*bookStoreShoppingCartItemVO.SellingPrice
			}
			if priceTotal < 1 {
				return errors.New("订单价格异常！"), orderNo
			}
			bookStoreOrder.CreateTime = common.JSONTime{Time: time.Now()}
			bookStoreOrder.UpdateTime = common.JSONTime{Time: time.Now()}
			bookStoreOrder.TotalPrice = priceTotal
			bookStoreOrder.ExtraInfo = ""
			bookStoreOrder.AddressId = userAddress.AddressId
			//生成订单项并保存订单项纪录
			if err = global.GVA_DB.Save(&bookStoreOrder).Error; err != nil {
				return errors.New("订单入库失败！"), orderNo
			}
			//生成订单收货地址快照，并保存至数据库
			var bookStoreOrderAddress mall.MallOrderAddress
			copier.Copy(&bookStoreOrderAddress, &userAddress)
			bookStoreOrderAddress.OrderId = bookStoreOrder.OrderId
			//生成所有的订单项快照，并保存至数据库
			var bookStoreOrderItems []manage.MallOrderItem
			for _, bookStoreShoppingCartItemVO := range myShoppingCartItems {
				var bookStoreOrderItem manage.MallOrderItem
				copier.Copy(&bookStoreOrderItem, &bookStoreShoppingCartItemVO)
				bookStoreOrderItem.OrderId = bookStoreOrder.OrderId
				bookStoreOrderItem.CreateTime = common.JSONTime{Time: time.Now()}
				bookStoreOrderItems = append(bookStoreOrderItems, bookStoreOrderItem)
			}
			if err = global.GVA_DB.Save(&bookStoreOrderItems).Error; err != nil {
				return err, orderNo
			}
		}
	}
	return
}

// PaySuccess 支付订单
func (m *MallOrderService) PaySuccess(orderNo string, payType int) (err error) {
	var mallOrder manage.MallOrder
	err = global.GVA_DB.Where("order_no = ? and is_deleted=0 ", orderNo).First(&mallOrder).Error
	if mallOrder != (manage.MallOrder{}) {
		if mallOrder.OrderStatus != 0 {
			return errors.New("订单状态异常！")
		}
		mallOrder.OrderStatus = enum.ORDER_PAID.Code()
		mallOrder.PayType = payType
		mallOrder.PayStatus = 1
		mallOrder.PayTime = common.JSONTime{time.Now()}
		mallOrder.UpdateTime = common.JSONTime{time.Now()}
		err = global.GVA_DB.Save(&mallOrder).Error
	}
	return
}

// FinishOrder 完结订单
func (m *MallOrderService) FinishOrder(token string, orderNo string) (err error) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}
	var mallOrder manage.MallOrder
	if err = global.GVA_DB.Where("order_no=? and is_deleted = 0", orderNo).First(&mallOrder).Error; err != nil {
		return errors.New("未查询到记录！")
	}
	if mallOrder.UserId != userToken.UserId {
		return errors.New("禁止该操作！")
	}
	mallOrder.OrderStatus = enum.ORDER_SUCCESS.Code()
	mallOrder.UpdateTime = common.JSONTime{time.Now()}
	err = global.GVA_DB.Save(&mallOrder).Error
	return
}

// CancelOrder 关闭订单
func (m *MallOrderService) CancelOrder(token string, orderNo string) (err error) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}
	var mallOrder manage.MallOrder
	if err = global.GVA_DB.Where("order_no=? and is_deleted = 0", orderNo).First(&mallOrder).Error; err != nil {
		return errors.New("未查询到记录！")
	}
	if mallOrder.UserId != userToken.UserId {
		return errors.New("禁止该操作！")
	}
	if utils.NumsInList(mallOrder.OrderStatus, []int{enum.ORDER_SUCCESS.Code(),
		enum.ORDER_CLOSED_BY_MALLUSER.Code(), enum.ORDER_CLOSED_BY_EXPIRED.Code(), enum.ORDER_CLOSED_BY_JUDGE.Code()}) {
		return errors.New("订单状态异常！")
	}
	mallOrder.OrderStatus = enum.ORDER_CLOSED_BY_MALLUSER.Code()
	mallOrder.UpdateTime = common.JSONTime{time.Now()}
	err = global.GVA_DB.Save(&mallOrder).Error
	return
}

// GetOrderDetailByOrderNo 获取订单详情
func (m *MallOrderService) GetOrderDetailByOrderNo(token string, orderNo string) (err error, orderDetail mallRes.MallOrderDetailVO) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), orderDetail
	}
	var mallOrder manage.MallOrder
	if err = global.GVA_DB.Where("order_no=? and is_deleted = 0", orderNo).First(&mallOrder).Error; err != nil {
		return errors.New("未查询到记录！"), orderDetail
	}
	if mallOrder.UserId != userToken.UserId {
		return errors.New("禁止该操作！"), orderDetail
	}
	var orderItems []manage.MallOrderItem
	err = global.GVA_DB.Where("order_id = ?", mallOrder.OrderId).Find(&orderItems).Error
	if len(orderItems) <= 0 {
		return errors.New("订单项不存在！"), orderDetail
	}

	var bookStoreOrderItemVOS []mallRes.BookStoreOrderItemVO
	copier.Copy(&bookStoreOrderItemVOS, &orderItems)
	copier.Copy(&orderDetail, &mallOrder)
	// 订单状态前端显示为中文
	_, OrderStatusStr := enum.GetBookStoreOrderStatusEnumByStatus(orderDetail.OrderStatus)
	_, payTapStr := enum.GetBookStoreOrderStatusEnumByStatus(orderDetail.PayType)
	orderDetail.OrderStatusString = OrderStatusStr
	orderDetail.PayTypeString = payTapStr
	orderDetail.BookStoreOrderItemVOS = bookStoreOrderItemVOS

	return
}

// MallOrderListBySearch 搜索订单
func (m *MallOrderService) MallOrderListBySearch(token string, pageNumber int, status string) (err error, list []mallRes.MallOrderResponse, total int64) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), list, total
	}
	// 根据搜索条件查询
	var bookStoreOrders []manage.MallOrder
	db := global.GVA_DB.Model(&bookStoreOrders)

	if status != "" {
		db.Where("order_status = ?", status)
	}
	err = db.Where("user_id =? and is_deleted=0 ", userToken.UserId).Count(&total).Error
	//这里前段没有做滚动加载，直接显示全部订单
	//limit := 5
	offset := 5 * (pageNumber - 1)
	err = db.Offset(offset).Order(" order_id desc").Find(&bookStoreOrders).Error

	var orderListVOS []mallRes.MallOrderResponse
	if total > 0 {
		//数据转换 将实体类转成vo
		copier.Copy(&orderListVOS, &bookStoreOrders)
		//设置订单状态中文显示值
		for _, bookStoreOrderListVO := range orderListVOS {
			_, statusStr := enum.GetBookStoreOrderStatusEnumByStatus(bookStoreOrderListVO.OrderStatus)
			bookStoreOrderListVO.OrderStatusString = statusStr
		}
		// 返回订单id
		var orderIds []int
		for _, order := range bookStoreOrders {
			orderIds = append(orderIds, order.OrderId)
		}
		//获取OrderItem
		var orderItems []manage.MallOrderItem
		if len(orderIds) > 0 {
			global.GVA_DB.Where("order_id in ?", orderIds).Find(&orderItems)
			itemByOrderIdMap := make(map[int][]manage.MallOrderItem)
			for _, orderItem := range orderItems {
				itemByOrderIdMap[orderItem.OrderId] = []manage.MallOrderItem{}
			}
			for k, v := range itemByOrderIdMap {
				for _, orderItem := range orderItems {
					if k == orderItem.OrderId {
						v = append(v, orderItem)
					}
					itemByOrderIdMap[k] = v
				}
			}
			//封装每个订单列表对象的订单项数据
			for _, bookStoreOrderListVO := range orderListVOS {
				if _, ok := itemByOrderIdMap[bookStoreOrderListVO.OrderId]; ok {
					orderItemListTemp := itemByOrderIdMap[bookStoreOrderListVO.OrderId]
					var bookStoreOrderItemVOS []mallRes.BookStoreOrderItemVO
					copier.Copy(&bookStoreOrderItemVOS, &orderItemListTemp)
					bookStoreOrderListVO.BookStoreOrderItemVOS = bookStoreOrderItemVOS
					_, OrderStatusStr := enum.GetBookStoreOrderStatusEnumByStatus(bookStoreOrderListVO.OrderStatus)
					bookStoreOrderListVO.OrderStatusString = OrderStatusStr
					list = append(list, bookStoreOrderListVO)
				}
			}
		}
	}
	return err, list, total
}
