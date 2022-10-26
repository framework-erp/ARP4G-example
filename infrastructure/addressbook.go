package infrastructure

import (
	"context"
	"example/aggregate"

	"github.com/zhengchengdong/ARP4G/arp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// type ContactRepository interface {
// 	Find(ctx context.Context, id any) (entity *aggregate.Contact, found bool)
// 	Take(ctx context.Context, id any) (entity *aggregate.Contact, found bool)
// 	Put(ctx context.Context, id any, entity *aggregate.Contact)
// 	Remove(ctx context.Context, id any) (removed *aggregate.Contact, exists bool)
// 	//模糊查找
// 	FindContains(ctx context.Context, contains string) ([]*aggregate.Contact, error)
// 	FindAllForGroup(ctx context.Context, groupId int64) ([]*aggregate.Contact, error)
// }

type ContactRepositoryImpl struct {
	arp.Repository[*aggregate.Contact]
	coll *mongo.Collection
}

func (repo *ContactRepositoryImpl) FindContains(ctx context.Context, contains string) ([]*aggregate.Contact, error) {
	repo.coll.Find(ctx, bson.M{"name": primitive.Regex{Pattern: "*" + contains + "*", Options: "im"}})
}
