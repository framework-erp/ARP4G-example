package repoimpl

import (
	"context"
	"example/aggregate"

	"github.com/framework-arp/ARP4G-mongodb/mongorepo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GroupRepositoryImpl struct {
	*mongorepo.MongodbRepository[*aggregate.Group]
	coll *mongo.Collection
}

func (repo *GroupRepositoryImpl) GetAll(ctx context.Context) ([]*aggregate.Group, error) {
	return repo.QueryAllByField(ctx, "state", 0)
}

func (repo *GroupRepositoryImpl) GetAllDeletedNotEmpty(ctx context.Context) ([]*aggregate.Group, error) {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"state", 1}},
				bson.D{{"contactcount", bson.D{{"$gt", 0}}}},
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

func NewGroupRepositoryImpl(mongoClient *mongo.Client) *GroupRepositoryImpl {
	mongoGroupRepository := mongorepo.NewMongodbRepository(mongoClient, "example", "Group", func() *aggregate.Group { return &aggregate.Group{} })
	var coll *mongo.Collection
	if mongoClient != nil {
		coll = mongoClient.Database("example").Collection("Group")
	}
	return &GroupRepositoryImpl{mongoGroupRepository, coll}
}
