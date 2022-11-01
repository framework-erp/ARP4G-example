# ARP4G-example
这是一个用ARP4G开发的简单的通讯录，用于给应用开发者演示ARP4G的功能。

## 安装

1. 首先需要 [Go](https://golang.org/) 已安装（**1.18及以上版本**），和[MongoDB](https://www.mongodb.com/try/download/community)，然后可以用以下命令安装ARP4G-example。

```sh
git clone https://github.com/framework-arp/ARP4G-example.git
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
你新添加了一个联系人。

想要运行更多的功能接口请查阅代码[gin 路由](https://github.com/framework-arp/ARP4G-example/blob/master/routers/addressbook.go)

建议从参考代码的角度，[AddressBookService](https://github.com/framework-arp/ARP4G-example/blob/master/service/addressbook.go)中的service定义是一个更好的入口：
```go
//通讯录的服务
type AddressBookService interface {
	//添加联系人
	AddContact(ctx context.Context, contactName string, phoneNumber string) (*aggregate.Contact, error)
	//删除联系人
	RemoveContact(ctx context.Context, contactId int64) error
	//把联系人放入组
	PutContactInGroup(ctx context.Context, contactId int64, groupId int64) (*aggregate.Contact, error)
	//添加组
	AddGroup(ctx context.Context, groupName string) (*aggregate.Group, error)
	//删除组，该组的联系人不会被删除
	RemoveGroup(ctx context.Context, groupId int64) error
	//获得所有的组
	GetGroups(ctx context.Context) ([]*aggregate.Group, error)
	//获得一个组的所有联系人
	GetContactsForGroup(ctx context.Context, groupId int64) ([]*aggregate.Contact, error)
	//获得不在任何组的所有联系人，这些联系人通常会在客户端被包装成一个“朋友”组
	GetContactsNotInGroup(ctx context.Context) ([]*aggregate.Contact, error)
	//模糊查询名字中包含特定文字的所有联系人
	QueryContacts(ctx context.Context, contains string) ([]*aggregate.Contact, error)
	FindContactById(ctx context.Context, contactId int64) (*aggregate.Contact, error)
}
```
