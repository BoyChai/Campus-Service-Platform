package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var Code map[string]string

func SendMsg(mobile string) error {
	// 随机数
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(9000) + 1000
	v := url.Values{}
	_now := strconv.FormatInt(time.Now().Unix(), 10)
	//fmt.Printf(_now)
	_account := "xxx"  //查看用户名 登录用户中心->验证码通知短信>产品总览->API接口信息->APIID
	_password := "xxx" //查看密码 登录用户中心->验证码通知短信>产品总览->API接口信息->APIKEY
	_content := fmt.Sprint("您的验证码是：", randomNumber, "。请不要把验证码泄露给其他人。")
	v.Set("account", _account)
	v.Set("password", _password)
	v.Set("mobile", mobile)
	v.Set("content", _content)
	v.Set("time", _now)
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://106.ihuyi.com/webservice/sms.php?method=Submit&format=json", body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//fmt.Printf("%+v\n", req) //看下发送的结构

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	if err != nil {
		fmt.Println("发送验证码错误：" + err.Error())
		return err
	}
	res, _ := ioutil.ReadAll(resp.Body)
	var data struct {
		Code  int    `json:"code"`
		Msg   string `json:"msg"`
		Smsid string `json:"smsid"`
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Println("发送验证码解析json出错:", err.Error())
		return errors.New("发送验证码解析json出错:" + err.Error())
	}
	if data.Code != 2 {
		fmt.Println("发送短信验证码出错:" + data.Msg)
		return errors.New(fmt.Sprint("发送短信验证码出错:" + data.Msg))
	}
	Code[mobile] = strconv.Itoa(randomNumber)
	return nil
}
