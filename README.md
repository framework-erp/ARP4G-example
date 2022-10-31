# ARP4G-example
这是一个用ARP4G开发的简单的通讯录，用于给应用开发者演示ARP4G的功能。

## 安装

1. 首先需要 [Go](https://golang.org/) 已安装（**1.18及以上版本**），和[MongoDB](https://www.mongodb.com/try/download/community)，然后可以用以下命令安装ARP4G-example。

```sh
git clone https://github.com/zhengchengdong/ARP4G-example.git
```

2. 运行ARP4G-example
```sh
cd ARP4G-example/cmd
go run main.go
```
这会启动一个HTTP服务在8080端口，看见以下提示说明启动成功：
```sh
[GIN-debug] Listening and serving HTTP on :8080
```

3. 功能验证


在浏览器输入：
```
http://127.0.0.1:8080/addressbook/addcontact?name=neo&phone=12346
```
你将看到返回的JSON：
```json
{
	"data": {
		"Id": 1586935438273155072,
		"Name": "neo",
		"PhoneNumber": "12346",
		"GroupId": 0
	},
	"success": true
}
```
你新添加了一个联系人

更多的功能接口请查阅[路由配置](https://github.com/zhengchengdong/ARP4G-example/blob/master/routers/addressbook.go)
