package arpservice

import (
	"context"
	"example/aggregate"
	"example/infrastructure/idgimpl"
	"example/infrastructure/repoimpl"
	"example/service"

	"github.com/bwmarrin/snowflake"
	"github.com/framework-arp/ARP4G/arp"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArpAddressBookService struct {
	service.AddressBookService
}

func (service *ArpAddressBookService) AddContact(ctx context.Context, contactName string, phoneNumber string) (*aggregate.Contact, error) {
	var contactToReturn *aggregate.Contact
	err := arp.Go(ctx, func(ctx context.Context) error {
		contact, err := service.AddressBookService.AddContact(ctx, contactName, phoneNumber)
		contactToReturn = contact
		return err
	})
	return contactToReturn, err
}

func (service *ArpAddressBookService) RemoveContact(ctx context.Context, contactId int64) error {
	return arp.Go(ctx, func(ctx context.Context) error {
		return service.AddressBookService.RemoveContact(ctx, contactId)
	})
}

func (service *ArpAddressBookService) PutContactInGroup(ctx context.Context, contactId int64, groupId int64) (*aggregate.Contact, error) {
	var contactToReturn *aggregate.Contact
	err := arp.Go(ctx, func(ctx context.Context) error {
		contact, err := service.AddressBookService.PutContactInGroup(ctx, contactId, groupId)
		contactToReturn = contact
		return err
	})
	return contactToReturn, err
}

func (service *ArpAddressBookService) AddGroup(ctx context.Context, groupName string) (*aggregate.Group, error) {
	var groupToReturn *aggregate.Group
	err := arp.Go(ctx, func(ctx context.Context) error {
		group, err := service.AddressBookService.AddGroup(ctx, groupName)
		groupToReturn = group
		return err
	})
	return groupToReturn, err
}

func (service *ArpAddressBookService) RemoveGroup(ctx context.Context, groupId int64) error {
	return arp.Go(ctx, func(ctx context.Context) error {
		return service.AddressBookService.RemoveGroup(ctx, groupId)
	})
}

func NewArpAddressBookService(mongoClient *mongo.Client, node *snowflake.Node) *ArpAddressBookService {
	addressBookServiceImpl := &service.AddressBookServiceImpl{repoimpl.NewContactRepositoryImpl(mongoClient), repoimpl.NewGroupRepositoryImpl(mongoClient),
		&idgimpl.SnowflakeContactIdGenerator{node}, &idgimpl.SnowflakeGroupIdGenerator{node}}
	return &ArpAddressBookService{addressBookServiceImpl}
}
