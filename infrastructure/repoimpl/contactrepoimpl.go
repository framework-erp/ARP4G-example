package repoimpl

import (
	"context"
	"example/aggregate"

	"github.com/framework-arp/ARP4G-mongodb/mongorepo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContactRepositoryImpl struct {
	*mongorepo.MongodbRepository[*aggregate.Contact]
	coll *mongo.Collection
}

func (repo *ContactRepositoryImpl) FindContains(ctx context.Context, contains string) ([]*aggregate.Contact, error) {
	cursor, err := repo.coll.Find(ctx, bson.M{"name": primitive.Regex{Pattern: contains, Options: "im"}})
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
	return repo.QueryAllByField(ctx, "groupid", groupId)
}

func NewContactRepositoryImpl(mongoClient *mongo.Client) *ContactRepositoryImpl {
	mongoContactRepository := mongorepo.NewMongodbRepository(mongoClient, "example", "Contact", func() *aggregate.Contact { return &aggregate.Contact{} })
	var coll *mongo.Collection
	if mongoClient != nil {
		coll = mongoClient.Database("example").Collection("Contact")
	}
	return &ContactRepositoryImpl{mongoContactRepository, coll}
}
