package infrastructure

import (
	"context"
	"example/aggregate"

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
