本项目由开源项目先锋商城修改而来，希望项目能更加完善，暂时前端只适配了移动端，后续有时间会适配其他
## 1. 俩前端项目

```
使用cnpm安装依赖，如果没有cnpm可以使用npm先安装cnpm和vite
cnpm install
使用vite启动和管理前端vue3项目
vite
正确启动会出现访问地址：
VITE v4.3.3  ready in 1223 ms
 ➜  Local:   http://localhost:8080/
前端项目的端口在vite.config.js文件中的server的port字段进行修改；
用户端front端口设置为了8080，admin管理端设置为了5173；5173是vite启动的项目默认端口
```
## 2. 后端

创建数据库bookstore

```
数据库名称选择：
bookstore
字符集和字符编码选择：
utf8mb4
utf8mb4_general_ci
```
然后跑SQL生成数据库初始数据： /bookstore-server/static-files/bookstore.xxx.sql
建议在navicat操作。
修改bookstore-server config.yaml中mysql的密码为你本地MySQL的密码
最后即可启动后端项目。
后端项目端口默认设置是8888，在config.yaml文件中system可以修改；

## 3. 账号密码

默认账号密码：
```
admin管理端：
admin
123456
```
## 4.项目简介
```
### 4.1. bookstore-front
用户端，包含用户的操作逻辑

src/router 前端路由
src/service 对应后端接口
src/views 页面实现

### 4.2. bookstore-admin
管理端，可以对商城的基础配置进行操作
### 4.3. bookstore-server
后端服务，前端的操作逻辑都在后端gin中实现，router >> handler >> service >> model ；
handler都在api/v1/目录下，service是具体实现都在service/目录下，model是数据库模型映射关系在model/目录下；
子目录mall/是用户端逻辑，manage/是管理端后端逻辑；
```

