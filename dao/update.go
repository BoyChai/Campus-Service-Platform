package dao

import (
	"errors"
	"fmt"
)

// ReceivingOrder 接单
func (d *dao) ReceivingOrder(orderID string, work uint) error {
	tx := d.db.Model(&WorkOrder{}).Where("id = ?", orderID).Updates(map[string]interface{}{
		"work_user":    work,
		"order_status": InProgress,
		"operator":     work,
	})
	if tx.Error != nil {
		fmt.Println("接单更新订单数据时产生错误", tx.Error)
		return errors.New(fmt.Sprint("接单更新订单数据时产生错误", tx.Error))
	}
	return nil
}

// CancellationOrder 取消
func (d *dao) CancellationOrder(orderID string, id uint) error {
	tx := d.db.Model(&WorkOrder{}).Where("id = ?", orderID).Updates(map[string]interface{}{
		"order_status": Cancellation,
		"operator":     id,
	})
	if tx.Error != nil {
		fmt.Println("取消订单时产生错误", tx.Error)
		return errors.New(fmt.Sprint("取消订单时产生错误", tx.Error))
	}
	return nil
}

// RemoveOrder 删除
func (d *dao) RemoveOrder(orderID string, id uint) error {
	tx := d.db.Model(&WorkOrder{}).Where("id = ?", orderID).Updates(map[string]interface{}{
		"operator": id,
	})
	if tx.Error != nil {
		fmt.Println("取消订单设置操作员时产生错误", tx.Error)
		return errors.New(fmt.Sprint("取消订单设置操作员时产生错误", tx.Error))
	}
	tx = d.db.Delete(&WorkOrder{}, orderID)
	if tx.Error != nil {
		fmt.Println("取消订单移除订单时产生错误", tx.Error)
		return errors.New(fmt.Sprint("取消订单移除订单时产生错误", tx.Error))
	}
	return nil
}
