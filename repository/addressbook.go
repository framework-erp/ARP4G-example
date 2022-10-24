package repository

import (
	"context"
	"example/aggregate"
)

type ContactRepository interface {
	Find(ctx context.Context, id any) (entity aggregate.Contact, found bool)
	Take(ctx context.Context, id any) (entity aggregate.Contact, found bool)
	Put(ctx context.Context, id any, entity aggregate.Contact)
	Remove(ctx context.Context, id any) (removed aggregate.Contact, exists bool)
	//模糊查找
	FindContains(ctx context.Context, contains string) ([]aggregate.Contact, error)
	FindAllForGroup(ctx context.Context, groupId int64) ([]aggregate.Contact, error)
}

type GroupRepository interface {
	Find(ctx context.Context, id any) (entity aggregate.Group, found bool)
	Take(ctx context.Context, id any) (entity aggregate.Group, found bool)
	Put(ctx context.Context, id any, entity aggregate.Group)
	GetAll(ctx context.Context) ([]aggregate.Group, error)
}
