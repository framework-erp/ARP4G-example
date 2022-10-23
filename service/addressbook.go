package service

import (
	"context"
	"example/aggregate"
)

//通讯录的服务
type AddressBookService interface {
	//添加联系人
	AddContact(ctx context.Context, contactName string)
	//删除联系人
	RemoveContact(ctx context.Context, contactId int64)
	//把联系人放入组
	PutContactInGroup(ctx context.Context, contactId int64, groupId int64)
	//添加组
	AddGroup(ctx context.Context, groupName string)
	//删除组，该组的联系人不会被删除
	RemoveGroup(ctx context.Context, groupId int64)
	//获得所有的组
	GetGroups(ctx context.Context) []aggregate.Group
	//获得一个组的所有联系人
	GetContactsForGroup(ctx context.Context, groupId int64) []aggregate.Contact
	//获得不在任何组的所有联系人，这些联系人通常会在客户端被包装成一个“朋友”组
	GetContactsNotInGroup(ctx context.Context, groupId int64) []aggregate.Contact
	//模糊查询名字中包含特定文字的所有联系人
	QueryContacts(ctx context.Context, contains string) []aggregate.Contact
}
