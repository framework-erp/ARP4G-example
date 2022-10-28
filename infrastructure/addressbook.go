package infrastructure

import (
	"context"
	"example/aggregate"
	"example/service"

	"github.com/zhengchengdong/ARP4G/arp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContactRepositoryImpl struct {
	arp.QueryRepository[*aggregate.Contact]
	coll *mongo.Collection
}

func (repo *ContactRepositoryImpl) FindContains(ctx context.Context, contains string) ([]*aggregate.Contact, error) {
	cursor, err := repo.coll.Find(ctx, bson.M{"name": primitive.Regex{Pattern: "*" + contains + "*", Options: "im"}})
	if err != nil {
		return nil, err
	}
	var results []bson.D
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	contacts := make([]*aggregate.Contact, 0)
	for _, result := range results {
		var doc []byte
		if doc, err = bson.Marshal(result); err != nil {
			return nil, err
		}
		contact := &aggregate.Contact{}
		bson.Unmarshal(doc, contact)
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (repo *ContactRepositoryImpl) FindAllForGroup(ctx context.Context, groupId int64) ([]*aggregate.Contact, error) {
	return repo.QueryAllByField(ctx, "GroupId", groupId)
}

type GroupRepositoryImpl struct {
	arp.QueryRepository[*aggregate.Group]
	coll *mongo.Collection
}

func (repo *GroupRepositoryImpl) GetAll(ctx context.Context) ([]*aggregate.Group, error) {
	return repo.QueryAllByField(ctx, "State", 0)
}

func (repo *GroupRepositoryImpl) GetAllDeletedNotEmpty(ctx context.Context) ([]*aggregate.Group, error) {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"State", 1}},
				bson.D{{"ContactCount", bson.D{{"$gt", 0}}}},
			},
		},
	}
	cursor, err := repo.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var results []bson.D
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	groups := make([]*aggregate.Group, 0)
	for _, result := range results {
		var doc []byte
		if doc, err = bson.Marshal(result); err != nil {
			return nil, err
		}
		group := &aggregate.Group{}
		bson.Unmarshal(doc, group)
		groups = append(groups, group)
	}
	return groups, nil
}

type ArpAddressBookService struct {
	service.AddressBookService
}

func (service *ArpAddressBookService) AddContact(ctx context.Context, contactName string, phoneNumber string) error {
	return arp.Go(ctx, func(ctx context.Context) error {
		return service.AddressBookService.AddContact(ctx, contactName, phoneNumber)
	})
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

func (service *ArpAddressBookService) AddGroup(ctx context.Context, groupName string) error {
	return arp.Go(ctx, func(ctx context.Context) error {
		return service.AddressBookService.AddGroup(ctx, groupName)
	})
}

func (service *ArpAddressBookService) RemoveGroup(ctx context.Context, groupId int64) error {
	return arp.Go(ctx, func(ctx context.Context) error {
		return service.AddressBookService.RemoveGroup(ctx, groupId)
	})
}
