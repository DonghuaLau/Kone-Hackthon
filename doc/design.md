
# 设计

## 数据库

### 访问
* user_id
* name
* create_time
* visit_date
* face_id

### 当前访问
* user_id
* last_visit_time
* is_visiting
* position_info

### 办公楼常驻成员
* user_id
* type
* name
* face_id


## 后台管理(web)
* /manager/index

### 管理员登录
* /manager/login

### 人脸采集
* 页面: /manager/face/collect
* 上传: /manager/face/upload
* 图像URL: /images/face/[id].jpg

### 访客授权
* 页面: /manager/guest/new
* 创建: /manager/guest/create
* 获取: /manager/guest/get

### 监控视频
* /manager/video/stream

## 用户界面

### 访客
* 访客信息，导航展示: /guest/[guest_id]
* 访客人脸采集: /guest/face/collect/[guest_id]

## 后台服务器 
* 人脸检测, 记录进入办公楼的访客

