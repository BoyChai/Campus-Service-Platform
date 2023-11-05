package dao

import (
	"errors"
	"fmt"
)

// CheckUser 查询用户是否合法,并返回用户数据
func (d *dao) CheckUser(number string, pass string) (User, error) {
	var u User
	tx := d.db.Where("number = ? and pass = ?", number, pass).First(&u)
	if tx.Error != nil {
		return User{}, errors.New(fmt.Sprint("用户检查出现错误:", tx.Error.Error()))
	}
	u.Pass = ""
	return u, nil
}

// GetUser 通过id获取某个用户的信息
func (d *dao) GetUser(id uint) (User, error) {
	var u User
	tx := d.db.Where("id = ?", id).First(&u)
	if tx != nil {
		return User{}, errors.New(fmt.Sprint("用户检查出现错误:", tx.Error.Error()))
	}
	return u.Clear(), nil
}

// GetWXUser 通过wx_id获取某个用户的信息
func (d *dao) GetWXUser(WxID string) (User, error) {
	var u User
	tx := d.db.Where("wx_id = ?", WxID).First(&u)
	if tx.Error != nil {
		return User{}, errors.New(fmt.Sprint("用户检查出现错误:", tx.Error.Error()))
	}
	return u.Clear(), nil
}

// QueryIDOrder 通过id查询订单
func (d *dao) QueryIDOrder(id uint) (WorkOrder, error) {
	var order WorkOrder
	tx := d.db.Where("id = ?", id).First(&order)
	if tx.Error != nil {
		return WorkOrder{}, errors.New(fmt.Sprint("通过id查询订单出现错误:", tx.Error.Error()))
	}
	return order, nil
}

// QueryUserOrder 通过userID查询订单
func (d *dao) QueryUserOrder(id uint) ([]WorkOrder, error) {
	var orders []WorkOrder
	tx := d.db.Where("create_user = ?", id).Find(&orders)
	if tx.Error != nil {
		return []WorkOrder{}, errors.New(fmt.Sprint("通过user查询订单出现错误:", tx.Error.Error()))
	}
	return orders, nil
}

// QueryOrderMsg 通过订单id查询聊天记录
func (d *dao) QueryOrderMsg(id uint) ([]Chat, error) {
	var chats []Chat
	tx := d.db.Where("order_id = ?", id).Find(&chats)
	if tx.Error != nil {
		return []Chat{}, errors.New(fmt.Sprint("通过订单查询聊天记录出现错误:", tx.Error.Error()))
	}
	return chats, nil
}
