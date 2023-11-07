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
