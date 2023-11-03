package dao

import (
	"errors"
	"fmt"
	"time"
)

// CreateUser 创建用户
func (d *dao) CreateUser(name string, number string, pass string, imgUrl string, wx string, role Role) error {
	user := User{
		Name:   name,
		Number: number,
		Pass:   pass,
		ImgUrl: imgUrl,
		WxID:   wx,
		Role:   role,
	}
	tx := d.db.Create(&user)
	if tx.Error != nil {
		return errors.New(fmt.Sprint("创建用户出现错误:", tx.Error.Error()))
	}
	return nil
}

// CreateOrder 创建订单
func (d *dao) CreateOrder(user uint, orderType OrderType) error {
	order := WorkOrder{
		CreateUser: user,
		OrderType:  orderType,
		Info:       JSON{},
	}
	tx := d.db.Create(&order)
	if tx.Error != nil {
		return errors.New(fmt.Sprint("创建订单出现错误:", tx.Error.Error()))
	}
	return nil
}

// SendMessage 提交聊天记录
func (d *dao) SendMessage(orderID uint, sender Role, msg JSON) error {
	chat := Chat{
		OrderID:  orderID,
		Sender:   sender,
		SendTime: time.Now(),
		Message:  msg,
	}
	tx := d.db.Create(&chat)
	if tx.Error != nil {
		return errors.New(fmt.Sprint("提交聊天信息出现错误:", tx.Error.Error()))
	}
	return nil
}
