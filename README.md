# 基于智能移动端的校园服务平台

# 需求场景

- 场景一：学生A发现自习室的门锁坏掉，学生A可以使用自己携带的智能手机，启动报修系统，选择相应的故障类型、具体所在地，并选择拍照提交给报修系统数据中心。

- 场景二：教工B发现教室的电脑和投影仪发生故障或操作问题，教工B可以使用自己携带的智能手机，启动报修系统，选择相应的故障类型、具体所在地，并可以选择拍照或者语音的形式提交给报修系统数据中心。在线的教室管理人员会选择直接通过手机语音通话帮助教工B或者直接维修，并在问题解决后将提交的任务设为以解决。教工B会选择是否对后勤的维修提出建议或评价。

- 场景三：主管C是一名设备维护部门的负责人，他的工作之一是负责排查学校可能发生的事故隐患，他手机安装有报修系统的智能移动端，某一天夜里，网络突发故障，造成校园网瘫痪，报修系统会检测到网络无法访问这一情况通过系统通知把这一情况及时通知到主管C的手机上。除此之外，主管C还可以通过智能移动端，查看通过后台数据中心分析得出的分析报告，内容包括统计最近的问题类型，可能的问题预测等。

- 场景四:借助于微信企业号(或公众号)平台,和学校统一身份认证平台融合,设计班级二维码点名系统。当上课时，通过多媒体系统显示二维码,学生通过扫码签到,并完成各种统计工作。也可以创意设计，开发其他校园应用系统。

# 平台设计

## 赛项需求

数据存储中心和角色（权限）管理配置系统、智能管理移动端和智能用户移动端。

## 兼容平台

赛项是智能移动端的比赛，主要还是手机端的兼容，实现的前端平台应该兼容小程序，安卓、IOS。小程序直接直接写，安卓和IOS可以直接采用网页套壳(目前网页构建移动端的框架暂未找到)。前端主要重心放在小程序这边。

## 分析设计

### 用户系统

用户系统应有用户管理，角色管理。

用户这块注册应采用学号、真实姓名、手机号来注册。

角色这块目前分为三种，管理员、教室管理员、维修工、普通用户。

### 工单系统

后台保修，提交问题的角色不限制，提交内容应有

故障类型、具体所在地、拍照、备注信息。

信息应包含

保修时间、完成时间。

完成时应该有照片。

当发起工单的时候应该发送通知教室管理员。通知形式采用短信。

### 聊天系统-第一次修改添加

提交工单后，维修师傅或管理员需要和发起工单的人进行了解情况，这里需要聊天系统，聊天系统每次发送消息应该包含订单id，发送者id，消息内容，消息类型。

### 管理系统

管理用户系统，工单报表，统计信息等待

管理系统用来管理用户的角色，工单报表的类型统计，完成时常统计。

### 告警系统

出现维修订单应该提醒,采用短信（后端处理）

### 班级管理系统(往后推)

提供签到点名，找个班级管理系统，在需求里面涉及较少，但是实际工作量比工单的要大的多，工单完善之后在考虑这个。

# 数据库设计

## 用户表

| id   | name   | number     | pass      | img_url  | wx_id  | role | other(JSON(map[string]insertface{}))   |
| ---- | ------ | ---------- | --------- | -------- | ------ | ---- | -------------------------------------- |
| id   | 用户名 | 用户手机号 | md5(pass) | 图片地址 | 微信id | 角色 | 存储一些其他信息，例如用户的上下级关系 |

角色(role)目前分为下面几种

- Admin  超级管理员(一般是开发者，拥有全部权限)
- Leader 领导(一般是维修工的管理者，拥有大部分权限)
- Work   工作者(一般是学校的修为人员，拥有小部分权限)
- Ordinary  普通用户(一般是提交问题的人员，拥有小部分权限)

## 工单表

| id       | create_user | order_type | order_status | CreatedAt | UpdatedAt    | DeletedAt    | CompleteAT   | work_user        | operator(Role)   | Info(JSON(map[string]insertface{})) |
| -------- | ----------- | ---------- | ------------ | --------- | ------------ | ------------ | ------------ | ---------------- | ---------------- | ----------------------------------- |
| 唯一标识 | 创建用户    | 订单类型   | 订单状态     | 创建时间  | 订单更新时间 | 订单删除时间 | 订单完成时间 | 接单的工作人员ID | 当前状态的操作员 | 创建时的一些初始信息                |

## 聊天表
| id       | order_id | sender | send_time | message(JSON(map[string]insertface{})) |
| -------- | -------- | ------ | --------- | -------------------------------------- |
| 唯一标识 | 工单id   | 发送者 | 时间      | 消息内容，消息类型，图片信息           |

发送者应该是是一个Role类型，超级管理员和工作者的领导都有权限查看聊天记录，查询的时候应该靠工单的id进行查询。消息的具体内容存储在一个json里面，内容，类型，发送的图片，图片地址。。

# 类型对照表

## Role角色

```go
const (
	// Admin 超级管理员 0
	Admin Role = iota
	// Leader 领导 1
	Leader
	// Work 工作者 2
	Work
	// Ordinary 普通用户 3
	Ordinary
)
```

## OrderType订单类型

```go
const (
    // Power 电力故障 0
    Power OrderType = iota
    // Network 网络故障 1
    Network
    // Water 水源故障 2 
    Water
    // HVAC 暖通空调故障 3
    HVAC
    // Device 设备故障 4
    Device
    // Construction 建筑设施 5
    Construction
    // SecuritySystem 安全系统 6
    SecuritySystem
    // CampusTransportation 校园交通 7
    CampusTransportation
    // Health 卫生 8
    Health
    // Other 其他 9
    Other
)
```

## OrderStatus订单状态

```go
const (
	// Pending 待处理 0
	Pending OrderStatus = iota
	// InProgress 处理中 1
	InProgress
	// WaitingConfirm 等待确认 2
	WaitingConfirm
	// Success 处理成功 3
	Success
	// Cancellation 订单被取消 4
	Cancellation
)
```

# 后端API设计

## 用户系统相关

### 用户注册(实现)

> POST /user/sendsms

number 手机号

> POST /user/signup

name 名字

number 手机号

pass 密码

img_url 图片地址

wx_id(可选) 微信ID

code 短信验证码

### 用户登录(实现)

> POST /user/login

number 手机号

pass 密码

### 用户信息修改

> PUT /user/update

name 名字

number 手机号

img 头像

## 工单系统相关

### 通过状态获取订单(实现)

> POST /order/get

status 状态类型(使用状态的id)

### 创建订单(实现)

> POST /order/create

addres 地点

img 照片

info 备注

### 工作人员接单(实现)

> POST /order/receiving

id 订单id

### 取消订单(实现)

> POST /order/cancellation

id 订单id

### 删除订单(实现)

> POST /order/delete

id 订单id

### 完成订单(实现)

> POST /order/complete

id 订单id

## 管理系统相关

### 用户角色修改

> POST /user/update

### 订单状态修改

> POST /order/update

## 聊天系统相关

### 获取聊天记录

> GET /message/get

id 订单id

### 发送消息(完成)

> POST /message/send

id 订单id

send_id 发送者id



# 第三方API接入

## 互亿无线

短信验证码服务，用来注册和登录

https://user.ihuyi.com/

## upyun

又拍云的云存储，用来存储头像，订单图片

