package utils

import (
	"errors"
	"fmt"
	"github.com/upyun/go-sdk/v3/upyun"
	"io"
)

var Upyun *upyun.UpYun

func InitUpyun() {
	Upyun = upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   "demo",
		Operator: "op",
		Password: "password",
	})
}

// PutOrderImg 上传订单图片
func PutOrderImg(fileName string, file io.Reader) error {
	err := Upyun.Put(&upyun.PutObjectConfig{
		Path:   "/campus/order/" + fileName,
		Reader: file,
	})
	if err != nil {
		fmt.Println("upyun文件上传失败:" + err.Error())
		return errors.New("upyun文件上传失败:" + err.Error())
	}
	return nil
}

// PutHeadImg 上传头像图片
func PutHeadImg(fileName string, file io.Reader) error {
	err := Upyun.Put(&upyun.PutObjectConfig{
		Path:   "/campus/head/" + fileName,
		Reader: file,
	})
	if err != nil {
		fmt.Println("upyun文件上传失败:" + err.Error())
		return errors.New("upyun文件上传失败:" + err.Error())
	}
	return nil
}
