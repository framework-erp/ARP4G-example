package service

import (
	"context"
	"example/aggregate"
	"example/idgenerator"
	"example/repository"
)

//通讯录的服务
type AddressBookService interface {
	//添加联系人
	AddContact(ctx context.Context, contactName string, phoneNumber string)
	//删除联系人
	RemoveContact(ctx context.Context, contactId int64)
	//把联系人放入组
	PutContactInGroup(ctx context.Context, contactId int64, groupId int64)
	//添加组
	AddGroup(ctx context.Context, groupName string)
	//删除组，该组的联系人不会被删除
	RemoveGroup(ctx context.Context, groupId int64)
	//获得所有的组
	GetGroups(ctx context.Context) []*aggregate.Group
	//获得一个组的所有联系人
	GetContactsForGroup(ctx context.Context, groupId int64) []*aggregate.Contact
	//获得不在任何组的所有联系人，这些联系人通常会在客户端被包装成一个“朋友”组
	GetContactsNotInGroup(ctx context.Context) []*aggregate.Contact
	//模糊查询名字中包含特定文字的所有联系人
	QueryContacts(ctx context.Context, contains string) []*aggregate.Contact
}

type AddressBookServiceImpl struct {
	ContactRepository  repository.ContactRepository
	GroupRepository    repository.GroupRepository
	ContactIdGenerator idgenerator.ContactIdGenerator
	GroupIdGenerator   idgenerator.GroupIdGenerator
}

func (service *AddressBookServiceImpl) AddContact(ctx context.Context, contactName string, phoneNumber string) {
	id := service.ContactIdGenerator.GenerateId(ctx)
	contact := &aggregate.Contact{id, contactName, phoneNumber, 0}
	service.ContactRepository.Put(ctx, id, contact)
}

func (service *AddressBookServiceImpl) RemoveContact(ctx context.Context, contactId int64) {
	service.ContactRepository.Remove(ctx, contactId)
}

func (service *AddressBookServiceImpl) PutContactInGroup(ctx context.Context, contactId int64, groupId int64) {
	contact, found := service.ContactRepository.Take(ctx, contactId)
	if !found {
		return
	}
	group, found := service.GroupRepository.Take(ctx, groupId)
	if !found {
		return
	}
	originalGroupId := contact.GroupId
	if originalGroupId != 0 {
		group, found := service.GroupRepository.Take(ctx, originalGroupId)
		if found {
			group.RemoveFrom()
			if group.IsDead() {
				service.GroupRepository.Remove(ctx, originalGroupId)
			}
		}
	}
	contact.GroupId = groupId
	group.AddTo()
}

func (service *AddressBookServiceImpl) AddGroup(ctx context.Context, groupName string) {
	id := service.GroupIdGenerator.GenerateId(ctx)
	group := &aggregate.Group{id, groupName, 0, 0}
	service.GroupRepository.Put(ctx, id, group)
}

func (service *AddressBookServiceImpl) RemoveGroup(ctx context.Context, groupId int64) {
	group, found := service.GroupRepository.Take(ctx, groupId)
	if !found {
		return
	}
	group.SetAsRemoved()
	if group.IsDead() {
		service.GroupRepository.Remove(ctx, groupId)
		return
	}
}

func (service *AddressBookServiceImpl) GetGroups(ctx context.Context) []*aggregate.Group {
	groups, err := service.GroupRepository.GetAll(ctx)
	if err != nil {
		return nil
	}
	return groups
}

func (service *AddressBookServiceImpl) GetContactsForGroup(ctx context.Context, groupId int64) []*aggregate.Contact {
	contacts, err := service.ContactRepository.FindAllForGroup(ctx, groupId)
	if err != nil {
		return nil
	}
	return contacts
}

func (service *AddressBookServiceImpl) GetContactsNotInGroup(ctx context.Context, groupId int64) []*aggregate.Contact {
	contactsNoGroup, err := service.ContactRepository.FindAllForGroup(ctx, 0)
	if err != nil {
		return nil
	}
	contactsNotInGroup := contactsNoGroup
	deletedNotEmptyGroup, err := service.GroupRepository.GetAllDeletedNotEmpty(ctx)
	if err != nil {
		return nil
	}
	if len(deletedNotEmptyGroup) == 0 {
		return contactsNoGroup
	}
	for _, group := range deletedNotEmptyGroup {
		contacts, err := service.ContactRepository.FindAllForGroup(ctx, group.Id)
		if err != nil {
			return nil
		}
		contactsNotInGroup = append(contactsNotInGroup, contacts...)
	}
	return contactsNotInGroup
}

func (service *AddressBookServiceImpl) QueryContacts(ctx context.Context, contains string) []*aggregate.Contact {
	contacts, err := service.ContactRepository.FindContains(ctx, contains)
	if err != nil {
		return nil
	}
	return contacts
}
