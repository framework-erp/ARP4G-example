package service

import (
	"context"
	"errors"
	"example/aggregate"
	"example/idgenerator"
	"example/repository"
)

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

type AddressBookServiceImpl struct {
	ContactRepository  repository.ContactRepository
	GroupRepository    repository.GroupRepository
	ContactIdGenerator idgenerator.ContactIdGenerator
	GroupIdGenerator   idgenerator.GroupIdGenerator
}

func (service *AddressBookServiceImpl) AddContact(ctx context.Context, contactName string, phoneNumber string) (*aggregate.Contact, error) {
	id := service.ContactIdGenerator.GenerateId(ctx)
	contact := &aggregate.Contact{id, contactName, phoneNumber, 0}
	service.ContactRepository.Put(ctx, id, contact)
	return contact, nil
}

func (service *AddressBookServiceImpl) RemoveContact(ctx context.Context, contactId int64) error {
	removed, exists := service.ContactRepository.Remove(ctx, contactId)
	if exists && removed.GroupId != 0 {
		group, found := service.GroupRepository.Take(ctx, removed.GroupId)
		if found {
			group.RemoveFrom()
			if group.IsDead() {
				service.GroupRepository.Remove(ctx, removed.GroupId)
			}
		}
	}
	return nil
}

func (service *AddressBookServiceImpl) PutContactInGroup(ctx context.Context, contactId int64, groupId int64) (*aggregate.Contact, error) {
	contact, found := service.ContactRepository.Take(ctx, contactId)
	if !found {
		return nil, errors.New("contact not found")
	}
	group, found := service.GroupRepository.Take(ctx, groupId)
	if !found {
		return nil, errors.New("group not found")
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
	return contact, nil
}

func (service *AddressBookServiceImpl) AddGroup(ctx context.Context, groupName string) (*aggregate.Group, error) {
	id := service.GroupIdGenerator.GenerateId(ctx)
	group := &aggregate.Group{id, groupName, 0, 0}
	service.GroupRepository.Put(ctx, id, group)
	return group, nil
}

func (service *AddressBookServiceImpl) RemoveGroup(ctx context.Context, groupId int64) error {
	group, found := service.GroupRepository.Take(ctx, groupId)
	if !found {
		return errors.New("group not found")
	}
	group.SetAsRemoved()
	if group.IsDead() {
		service.GroupRepository.Remove(ctx, groupId)
	}
	return nil
}

func (service *AddressBookServiceImpl) GetGroups(ctx context.Context) ([]*aggregate.Group, error) {
	groups, err := service.GroupRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (service *AddressBookServiceImpl) GetContactsForGroup(ctx context.Context, groupId int64) ([]*aggregate.Contact, error) {
	contacts, err := service.ContactRepository.FindAllForGroup(ctx, groupId)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func (service *AddressBookServiceImpl) GetContactsNotInGroup(ctx context.Context) ([]*aggregate.Contact, error) {
	contactsNoGroup, err := service.ContactRepository.FindAllForGroup(ctx, 0)
	if err != nil {
		return nil, err
	}
	contactsNotInGroup := contactsNoGroup
	deletedNotEmptyGroup, err := service.GroupRepository.GetAllDeletedNotEmpty(ctx)
	if err != nil {
		return nil, err
	}
	if len(deletedNotEmptyGroup) == 0 {
		return contactsNoGroup, nil
	}
	for _, group := range deletedNotEmptyGroup {
		contacts, err := service.ContactRepository.FindAllForGroup(ctx, group.Id)
		if err != nil {
			return nil, err
		}
		contactsNotInGroup = append(contactsNotInGroup, contacts...)
	}
	return contactsNotInGroup, nil
}

func (service *AddressBookServiceImpl) QueryContacts(ctx context.Context, contains string) ([]*aggregate.Contact, error) {
	contacts, err := service.ContactRepository.FindContains(ctx, contains)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func (service *AddressBookServiceImpl) FindContactById(ctx context.Context, contactId int64) (*aggregate.Contact, error) {
	contact, found := service.ContactRepository.Find(ctx, contactId)
	if !found {
		return nil, errors.New("contact not found")
	}
	return contact, nil
}
